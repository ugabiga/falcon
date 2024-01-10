package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/ugabiga/falcon/pkg/config"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	UserName = "falcon"
)

type InfraStackProps struct {
	awscdk.StackProps
}

func NewStack(scope constructs.Construct, id string, cfg *config.Config, props *InfraStackProps) awscdk.Stack {
	var stackProps awscdk.StackProps
	if props != nil {
		stackProps = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &stackProps)

	u := newUser(stack)
	ecr := newECRRepository(stack)
	newUserPolicy(stack, u, ecr)

	vpc := lookupVPC(stack)
	vpSubnets := lookupVPCSubnets(stack)
	lambdaSecurityGroup := newLambdaSecurityGroup(stack, vpc)

	environment := newLambdaEnvironment(cfg)
	lambdaServerFunc := newLambdaServer(stack, ecr, environment)
	lambdaCronFunc := newLambdaCron(stack, ecr, environment)
	lambdaWorkerFunc := newLambdaWorker(stack, ecr, vpc, vpSubnets, lambdaSecurityGroup, environment)
	newLambdaPolicy(stack, lambdaServerFunc, lambdaCronFunc, lambdaWorkerFunc)

	return stack
}

func newLambdaPolicy(stack awscdk.Stack, lambdaServerFunc awslambda.DockerImageFunction, lambdaCronFunc awslambda.DockerImageFunction, lambdaWorkerFunc awslambda.DockerImageFunction) {
	lambdaPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("sqs:ChangeMessageVisibility"),
			jsii.String("sqs:DeleteMessage"),
			jsii.String("sqs:SendMessage"),
			jsii.String("sqs:GetQueueAttributes"),
			jsii.String("sqs:GetQueueUrl"),
			jsii.String("sqs:ReceiveMessage"),
		},
	})
	lambdaPolicy.AddAllResources()

	lambdaServerFunc.AddToRolePolicy(lambdaPolicy)
	lambdaCronFunc.AddToRolePolicy(lambdaPolicy)
	lambdaWorkerFunc.AddToRolePolicy(lambdaPolicy)

	lambdaDynamoPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("dynamodb:BatchGetItem"),
			jsii.String("dynamodb:BatchWriteItem"),
			jsii.String("dynamodb:DeleteItem"),
			jsii.String("dynamodb:GetItem"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:UpdateItem"),
		},
	})
	lambdaDynamoPolicy.AddResources(
		jsii.String("arn:aws:dynamodb:ap-northeast-2:358059338173:table/falcon"),
		jsii.String("arn:aws:dynamodb:ap-northeast-2:358059338173:table/falcon/index/*"),
	)
	lambdaServerFunc.AddToRolePolicy(lambdaDynamoPolicy)
	lambdaCronFunc.AddToRolePolicy(lambdaDynamoPolicy)
	lambdaWorkerFunc.AddToRolePolicy(lambdaDynamoPolicy)
}

func newLambdaWorker(stack awscdk.Stack, ecr awsecr.Repository, vpc awsec2.IVpc, vpcSubnets *awsec2.SubnetSelection, securityGroup awsec2.SecurityGroup, environment map[string]*string) awslambda.DockerImageFunction {
	lambdaWorkerName := "falcon-worker"
	workerSQSName := "falcon-worker-sqs"
	lambdaWorkerCmd := "worker"
	lambdaWorkerFunc := awslambda.NewDockerImageFunction(stack, jsii.String(lambdaWorkerName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String(lambdaWorkerCmd)},
		}),
		Architecture:      awslambda.Architecture_ARM_64(),
		Timeout:           awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention:      awslogs.RetentionDays_FIVE_DAYS,
		AllowPublicSubnet: jsii.Bool(true),
		Vpc:               vpc,
		VpcSubnets:        vpcSubnets,
		SecurityGroups:    &[]awsec2.ISecurityGroup{securityGroup},
		Environment:       &environment,
	})
	workerSQS := awssqs.NewQueue(stack, jsii.String("WorkerSQS"), &awssqs.QueueProps{
		QueueName:         jsii.String(workerSQSName),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(500)),
	})
	eventSource := awslambdaeventsources.NewSqsEventSource(workerSQS, &awslambdaeventsources.SqsEventSourceProps{})
	lambdaWorkerFunc.AddEventSource(eventSource)

	awscdk.NewCfnOutput(stack, jsii.String("lambdaWorkerName"), &awscdk.CfnOutputProps{
		Value: lambdaWorkerFunc.FunctionName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("sqsQueueName"), &awscdk.CfnOutputProps{
		Value: workerSQS.QueueName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("sqsQueueURL"), &awscdk.CfnOutputProps{
		Value: workerSQS.QueueUrl(),
	})

	return lambdaWorkerFunc
}

func newLambdaCron(stack awscdk.Stack, ecr awsecr.Repository, environment map[string]*string) awslambda.DockerImageFunction {
	lambdaCronName := "falcon-cron"
	lambdaCronRuleName := "falcon-cron-rule"
	lambdaCronCmd := "cron"
	lambdaCronFunc := awslambda.NewDockerImageFunction(stack, jsii.String(lambdaCronName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String(lambdaCronCmd)},
		}),
		Architecture: awslambda.Architecture_ARM_64(),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention: awslogs.RetentionDays_FIVE_DAYS,
		Environment:  &environment,
	})
	cronRule := awsevents.NewRule(stack, jsii.String(lambdaCronRuleName), &awsevents.RuleProps{
		//Schedule every 1 hour
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Minute: jsii.String("0"),
		}),
		////Schedule every 5 minute
		//Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
		//	Minute: jsii.String("*/5"),
		//}),
	})
	cronRule.AddTarget(awseventstargets.NewLambdaFunction(lambdaCronFunc, nil))

	awscdk.NewCfnOutput(stack, jsii.String("lambdaCronFuncName"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaCronFuncName"),
		Value:      lambdaCronFunc.FunctionName(),
	})

	return lambdaCronFunc
}

func newLambdaServer(stack awscdk.Stack, ecr awsecr.Repository, environment map[string]*string) awslambda.DockerImageFunction {
	publicDomainName := "api-falcon.vultor.xyz"
	publicDomainCertificateArn := "arn:aws:acm:us-east-1:358059338173:certificate/e74a2c12-794d-4ae4-849b-977baadf9965"
	lambdaServerName := "falcon-server"
	lambdaServerCmd := "lambda-server"

	lambdaServerFunc := awslambda.NewDockerImageFunction(stack, jsii.String(lambdaServerName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String(lambdaServerCmd)},
		}),
		Architecture: awslambda.Architecture_ARM_64(),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention: awslogs.RetentionDays_FIVE_DAYS,
		Environment:  &environment,
	})
	lambdaServerFunc.AddAlias(jsii.String("current"), &awslambda.AliasOptions{
		ProvisionedConcurrentExecutions: jsii.Number(1),
	})
	lambdaURL := lambdaServerFunc.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})
	// Add a CloudFront distribution to route between the public directory and the Lambda function URL.
	lambdaURLDomain := awscdk.Fn_Select(jsii.Number(2), awscdk.Fn_Split(jsii.String("/"), lambdaURL.Url(), nil))
	lambdaOrigin := awscloudfrontorigins.NewHttpOrigin(lambdaURLDomain, &awscloudfrontorigins.HttpOriginProps{
		ProtocolPolicy: awscloudfront.OriginProtocolPolicy_HTTPS_ONLY,
	})

	cf := awscloudfront.NewDistribution(stack, jsii.String("ServerCloudFrontDistribution"), &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			AllowedMethods:       awscloudfront.AllowedMethods_ALLOW_ALL(),
			Origin:               lambdaOrigin,
			CachedMethods:        awscloudfront.CachedMethods_CACHE_GET_HEAD(),
			OriginRequestPolicy:  awscloudfront.OriginRequestPolicy_ALL_VIEWER_EXCEPT_HOST_HEADER(),
			CachePolicy:          awscloudfront.CachePolicy_CACHING_DISABLED(),
			ViewerProtocolPolicy: awscloudfront.ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
		},
		Certificate: awscertificatemanager.Certificate_FromCertificateArn(stack,
			jsii.String("Certificate"),
			jsii.String(publicDomainCertificateArn),
		),
		DomainNames: &[]*string{
			jsii.String(publicDomainName),
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String("lambdaServerFuncName"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaServerFuncName"),
		Value:      lambdaServerFunc.FunctionName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionUrl"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaFunctionUrl"),
		Value:      lambdaURL.Url(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("cloudFrontDomainName"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("cloudFrontDomainName"),
		Value:      cf.DomainName(),
	})

	return lambdaServerFunc
}

func newLambdaEnvironment(cfg *config.Config) map[string]*string {
	//boot to string
	dynamoIsLocalStr := "false"
	if cfg.DynamoIsLocal {
		dynamoIsLocalStr = "true"
	}

	return map[string]*string{
		//"AWS_REGION":         do not use this aws provider will automatically set this value
		"DB_DRIVER_NAME":       jsii.String(cfg.DBDriverName),
		"DB_SOURCE":            jsii.String(cfg.DBSource),
		"GOOGLE_CLIENT_ID":     jsii.String(cfg.GoogleClientID),
		"GOOGLE_CLIENT_SECRET": jsii.String(cfg.GoogleClientSecret),
		"SESSION_SECRET_KEY":   jsii.String(cfg.SessionSecretKey),
		"JWT_SECRET_KEY":       jsii.String(cfg.JWTSecretKey),
		"WEB_URL":              jsii.String(cfg.WebURL),
		"ENCRYPTION_KEY":       jsii.String(cfg.EncryptionKey),
		"DYNAMO_IS_LOCAL":      jsii.String(dynamoIsLocalStr),
		"MESSAGING_PLATFORM":   jsii.String(cfg.MessagingPlatform),
		"SQS_QUEUE_URL":        jsii.String(cfg.SQSQueueURL),
	}
}

func newLambdaSecurityGroup(stack awscdk.Stack, vpc awsec2.IVpc) awsec2.SecurityGroup {
	lambdaSecurityGroup := "falcon-lambda-security-group"

	securityGroup := awsec2.NewSecurityGroup(stack, jsii.String("LambdaSecurityGroup"), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String(*stack.StackName() + "-" + lambdaSecurityGroup),
		AllowAllOutbound:  jsii.Bool(true),
		Description:       jsii.String("Allow Lambda functions to access resources"),
	})

	securityGroup.AddIngressRule(
		securityGroup,
		awsec2.NewPort(&awsec2.PortProps{
			Protocol:             awsec2.Protocol_ALL,
			StringRepresentation: jsii.String("Self Traffic"),
		}),
		jsii.String("Allow requests"),
		jsii.Bool(false),
	)
	return securityGroup
}

func lookupVPC(stack awscdk.Stack) awsec2.IVpc {
	VpcID := "vpc-07fe0ad3ccddd451c"
	return awsec2.Vpc_FromLookup(stack, jsii.String("DefaultVPC"), &awsec2.VpcLookupOptions{
		VpcId: jsii.String(VpcID),
	})
}

func lookupVPCSubnets(stack awscdk.Stack) *awsec2.SubnetSelection {
	// Use Private Subnets
	subNet1 := "subnet-05a0867a2d76c57d1"
	subNet2 := "subnet-08af91935c2a2c89a"

	return &awsec2.SubnetSelection{
		SubnetFilters: &[]awsec2.SubnetFilter{
			awsec2.SubnetFilter_ByIds(&[]*string{
				jsii.String(subNet1),
				jsii.String(subNet2),
			}),
		},
	}
}

func newUserPolicy(stack awscdk.Stack, user awsiam.User, ecr awsecr.Repository) []awsiam.Policy {
	ecrPolicyName := "falcon-ecr-policy"
	ecrServicePolicyName := "falcon-ecr-service-policy"
	deployLambdaPolicyName := "falcon-deploy-lambda-policy"

	ecrPolicy := awsiam.NewPolicy(stack, jsii.String("ECRPolicy"), &awsiam.PolicyProps{
		PolicyName: jsii.String(ecrPolicyName),
		Statements: &[]awsiam.PolicyStatement{
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Effect: awsiam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String("ecr:BatchGetImage"),
					jsii.String("ecr:BatchCheckLayerAvailability"),
					jsii.String("ecr:CompleteLayerUpload"),
					jsii.String("ecr:GetDownloadUrlForLayer"),
					jsii.String("ecr:InitiateLayerUpload"),
					jsii.String("ecr:PutImage"),
					jsii.String("ecr:UploadLayerPart"),
				},
				Resources: &[]*string{
					ecr.RepositoryArn(),
				},
			}),
		},
	})

	ecrServicePolicy := awsiam.NewPolicy(stack, jsii.String("ECRServicePolicy"), &awsiam.PolicyProps{
		PolicyName: jsii.String(ecrServicePolicyName),
		Statements: &[]awsiam.PolicyStatement{
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Effect: awsiam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String("ecr:GetAuthorizationToken"),
				},
				Resources: &[]*string{
					jsii.String("*"),
				},
			}),
		},
	})

	deployLambdaPolicy := awsiam.NewPolicy(stack, jsii.String("DeployLambdaPolicy"), &awsiam.PolicyProps{
		PolicyName: jsii.String(deployLambdaPolicyName),
		Statements: &[]awsiam.PolicyStatement{
			awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
				Effect: awsiam.Effect_ALLOW,
				Actions: &[]*string{
					jsii.String("lambda:CreateFunction"),
					jsii.String("lambda:UpdateFunctionCode"),
					jsii.String("lambda:UpdateFunctionConfiguration"),
					jsii.String("lambda:PublishVersion"),
					jsii.String("lambda:CreateAlias"),
					jsii.String("lambda:DeleteFunction"),
				},
				Resources: &[]*string{
					jsii.String("*"),
				},
			}),
		},
	})

	policies := []awsiam.Policy{ecrPolicy, ecrServicePolicy, deployLambdaPolicy}
	for _, policy := range policies {
		policy.AttachToUser(user)
	}

	return policies
}

func newECRRepository(stack awscdk.Stack) awsecr.Repository {
	ECRName := "falcon-ecr-repository"
	ecr := awsecr.NewRepository(stack, jsii.String("ECRRepository"), &awsecr.RepositoryProps{
		RepositoryName:     jsii.String(ECRName),
		RemovalPolicy:      awscdk.RemovalPolicy_DESTROY,
		ImageTagMutability: awsecr.TagMutability_MUTABLE,
		ImageScanOnPush:    jsii.Bool(false),
	})

	awscdk.NewCfnOutput(stack, jsii.String("ECRRepositoryName"), &awscdk.CfnOutputProps{
		Value: ecr.RepositoryName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("ECRRepositoryURI"), &awscdk.CfnOutputProps{
		Value: ecr.RepositoryUri(),
	})
	return ecr
}

func newUser(stack awscdk.Stack) awsiam.User {
	userName := "falcon-user"
	user := awsiam.NewUser(stack, jsii.String("User"), &awsiam.UserProps{
		UserName: jsii.String(userName),
	})
	awscdk.NewCfnOutput(stack, jsii.String("UserName"), &awscdk.CfnOutputProps{
		Value: user.UserName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("UserArn"), &awscdk.CfnOutputProps{
		Value: user.UserArn(),
	})

	accessKey := awsiam.NewAccessKey(stack, jsii.String("UserAccessKey"), &awsiam.AccessKeyProps{
		User: user,
	})
	awscdk.NewCfnOutput(stack, jsii.String("AccessKeyId"), &awscdk.CfnOutputProps{
		Value: accessKey.AccessKeyId(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("AccessKeySecret"), &awscdk.CfnOutputProps{
		Value: accessKey.SecretAccessKey().UnsafeUnwrap(),
	})

	return user
}

func newConfig() (*config.Config, error) {
	cfg := config.NewConfigWithoutSetting()
	if err := cfg.Load(
		&[]string{"../"}[0],
		&[]string{"config.prod"}[0],
	); err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	defer jsii.Close()

	cfg, err := newConfig()
	if err != nil {
		panic(err)
	}
	app := awscdk.NewApp(nil)

	NewStack(app, "FalconStack", cfg,
		&InfraStackProps{
			awscdk.StackProps{
				Env: env(),
			},
		})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("358059338173"),
		Region:  jsii.String("ap-northeast-2"),
	}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
