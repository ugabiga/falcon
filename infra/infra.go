package main

import (
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecspatterns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsrds"
	secretmgr "github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/ugabiga/falcon/pkg/config"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	UserName   = "falcon"
	ECRName    = "falcon"
	LambdaName = "falcon"
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
	//newNetworks(stack, vpc)
	lambdaSecurityGroup := newLambdaSecurityGroup(stack, vpc)
	newLambda(stack, ecr, cfg, vpc, lambdaSecurityGroup)
	//newLambda(stack, ecr, cfg, vpc)

	databaseSecurityGroup := newDatabaseSecurityGroup(stack, vpc, lambdaSecurityGroup)
	newDatabaseCluster(stack, vpc, databaseSecurityGroup)
	//newECSCluster(stack, cfg, vpc, ecr)

	return stack
}

func newNetworks(stack awscdk.Stack, vpc awsec2.IVpc) {
	//privateSubnetName := "falcon-private-subnet"
	//publicSubnetName := "falcon-public-subnet"
	natGatewayName := "falcon-nat-gateway"
	elasticIPName := "falcon-elastic-ip"

	//privateSubnet := awsec2.NewSubnet(stack, jsii.String(privateSubnetName), &awsec2.SubnetProps{
	//	CidrBlock:        jsii.String("10.0.0.0/16"),
	//	AvailabilityZone: jsii.String("ap-northeast-2a"),
	//	VpcId:            vpc.VpcId(),
	//})

	//publicSubnet := awsec2.NewSubnet(stack, jsii.String(publicSubnetName), &awsec2.SubnetProps{
	//	CidrBlock:        jsii.String("10.0.0.0/24"),
	//	AvailabilityZone: jsii.String("ap-northeast-2a"),
	//	VpcId:            vpc.VpcId(),
	//})

	subnet := awsec2.Subnet_FromSubnetId(
		stack, jsii.String("falcon-subnet"), jsii.String("subnet-0fc09c378f7e23e85"),
	)

	elasticIP := awsec2.NewCfnEIP(stack, jsii.String(elasticIPName), &awsec2.CfnEIPProps{
		Domain: jsii.String("vpc"),
	})

	awsec2.NewCfnNatGateway(stack, jsii.String(natGatewayName), &awsec2.CfnNatGatewayProps{
		SubnetId:         subnet.SubnetId(),
		AllocationId:     elasticIP.AttrAllocationId(),
		ConnectivityType: jsii.String("public"),
		Tags: &[]*awscdk.CfnTag{
			{
				Key:   jsii.String("Name"),
				Value: jsii.String("falcon-nat-gateway"),
			},
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String("ElasticIP"), &awscdk.CfnOutputProps{
		Value: elasticIP.AttrAllocationId(),
	})

}

func lookupVPC(stack awscdk.Stack) awsec2.IVpc {
	// Import default VPC.
	return awsec2.Vpc_FromLookup(stack, jsii.String("DefaultVPC"), &awsec2.VpcLookupOptions{
		VpcId: jsii.String("vpc-06b43122d657875bc"),
	})
}

func newECSCluster(stack awscdk.Stack, cfg *config.Config, vpc awsec2.IVpc, imageRepository awsecr.Repository) {
	clusterName := "falcon-ecs-cluster"
	loadBalancerName := "falcon-ecs-service-lb"

	cluster := awsecs.NewCluster(stack, jsii.String(clusterName), &awsecs.ClusterProps{
		Vpc: vpc,
	})

	awsecspatterns.NewApplicationLoadBalancedFargateService(stack, jsii.String("falcon-ecs-service"),
		&awsecspatterns.ApplicationLoadBalancedFargateServiceProps{
			Cluster:        cluster,
			DesiredCount:   jsii.Number(1),
			Cpu:            jsii.Number(512),
			MemoryLimitMiB: jsii.Number(1024),
			TaskImageOptions: &awsecspatterns.ApplicationLoadBalancedTaskImageOptions{
				Image:         awsecs.EcrImage_FromEcrRepository(imageRepository, jsii.String("latest")),
				ContainerPort: jsii.Number(8080),
				Environment: &map[string]*string{
					"DB_DRIVER_NAME":       jsii.String(cfg.DBDriverName),
					"DB_SOURCE":            jsii.String(cfg.DBSource),
					"SESSION_SECRET_KEY":   jsii.String(cfg.SessionSecretKey),
					"JWT_SECRET_KEY":       jsii.String(cfg.JWTSecretKey),
					"GOOGLE_CLIENT_ID":     jsii.String(cfg.GoogleClientID),
					"GOOGLE_CLIENT_SECRET": jsii.String(cfg.GoogleClientSecret),
					"WEB_URL":              jsii.String(cfg.WebURL),
					"ENCRYPTION_KEY":       jsii.String(cfg.EncryptionKey),
				},
				EnableLogging: jsii.Bool(true),
				LogDriver: awsecs.AwsLogDriver_AwsLogs(&awsecs.AwsLogDriverProps{
					StreamPrefix: jsii.String("falcon"),
				}),
			},
			AssignPublicIp:     jsii.Bool(true),
			LoadBalancerName:   jsii.String(loadBalancerName),
			PublicLoadBalancer: jsii.Bool(true),
		},
	)
}

func newLambdaSecurityGroup(stack awscdk.Stack, vpc awsec2.IVpc) awsec2.SecurityGroup {
	lambdaSecurityGroup := "falcon-lambda-security-group"

	// Create PostgreSQL Security Group.
	securityGroup := awsec2.NewSecurityGroup(stack, jsii.String(lambdaSecurityGroup), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String(*stack.StackName() + "-" + lambdaSecurityGroup),
		AllowAllOutbound:  jsii.Bool(true),
		Description:       jsii.String("PostgreSQL Security Group"),
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

func newDatabaseSecurityGroup(stack awscdk.Stack, vpc awsec2.IVpc, lambdaSecurityGroup awsec2.SecurityGroup) awsec2.SecurityGroup {
	myIP := "59.12.246.85/32"
	rdsSecurityGroup := "falcon-postgresql-security-group"

	// Create PostgreSQL Security Group.
	securityGroup := awsec2.NewSecurityGroup(stack, jsii.String(rdsSecurityGroup), &awsec2.SecurityGroupProps{
		Vpc:               vpc,
		SecurityGroupName: jsii.String(*stack.StackName() + "-" + rdsSecurityGroup),
		AllowAllOutbound:  jsii.Bool(true),
		Description:       jsii.String("PostgreSQL Security Group"),
	})

	securityGroup.AddIngressRule(
		awsec2.Peer_Ipv4(jsii.String(myIP)),
		awsec2.NewPort(&awsec2.PortProps{
			Protocol:             awsec2.Protocol_TCP,
			FromPort:             jsii.Number(5432),
			ToPort:               jsii.Number(5432),
			StringRepresentation: jsii.String("Standard Postgres"),
		}),
		jsii.String("Allow requests to Postgres DB instance."),
		jsii.Bool(false),
	)

	securityGroup.AddIngressRule(
		lambdaSecurityGroup,
		awsec2.NewPort(&awsec2.PortProps{
			Protocol:             awsec2.Protocol_TCP,
			FromPort:             jsii.Number(5432),
			ToPort:               jsii.Number(5432),
			StringRepresentation: jsii.String("Standard Postgres"),
		}),
		jsii.String("Allow requests to Postgres DB instance."),
		jsii.Bool(false),
	)

	return securityGroup
}

func newDatabaseCluster(stack awscdk.Stack, vpc awsec2.IVpc, securityGroup awsec2.SecurityGroup) {
	parameterGroupName := "falcon-postgresql-parameter-group"
	auroraClusterName := "falcon-postgresql-cluster"
	subnetGroupName := "falcon-subnet-group"
	rdsUserName := "falcon_admin"
	rdsDBName := "falcon"

	subnetGrp := awsrds.NewSubnetGroup(stack, jsii.String(subnetGroupName), &awsrds.SubnetGroupProps{
		Vpc:             vpc,
		RemovalPolicy:   awscdk.RemovalPolicy_DESTROY,
		SubnetGroupName: jsii.String(*stack.StackName() + "-" + subnetGroupName),
		VpcSubnets:      &awsec2.SubnetSelection{SubnetType: awsec2.SubnetType_PUBLIC},
		Description:     jsii.String("RDS SubnetGroup"),
	})

	engine := awsrds.DatabaseClusterEngine_AuroraPostgres(&awsrds.AuroraPostgresClusterEngineProps{
		Version: awsrds.AuroraPostgresEngineVersion_VER_15_2(),
	})

	parameterGroup := awsrds.NewParameterGroup(stack, jsii.String(parameterGroupName), &awsrds.ParameterGroupProps{
		Engine:      engine,
		Description: jsii.String("falcon RDS ParameterGroup"),
		Parameters:  &map[string]*string{},
	})

	secret := secretmgr.NewSecret(stack, jsii.String("falcon-postgresql-secret"), &secretmgr.SecretProps{
		SecretName: jsii.String("falcon-postgresql-secret"),
		GenerateSecretString: &secretmgr.SecretStringGenerator{
			SecretStringTemplate: jsii.String(fmt.Sprintf(`{"username": "%s"}`, rdsUserName)),
			ExcludePunctuation:   jsii.Bool(true),
			IncludeSpace:         jsii.Bool(false),
			GenerateStringKey:    jsii.String("password"),
		},
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	databaseCluster := awsrds.NewDatabaseCluster(stack, jsii.String(auroraClusterName), &awsrds.DatabaseClusterProps{
		Engine: engine,
		Writer: awsrds.ClusterInstance_ServerlessV2(jsii.String("writer"), &awsrds.ServerlessV2ClusterInstanceProps{
			PubliclyAccessible:       jsii.Bool(true),
			AllowMajorVersionUpgrade: jsii.Bool(false),
			AutoMinorVersionUpgrade:  jsii.Bool(true),
		}),
		ParameterGroup:      parameterGroup,
		DefaultDatabaseName: jsii.String(rdsDBName),
		Credentials:         awsrds.Credentials_FromSecret(secret, jsii.String(rdsUserName)),
		Backup: &awsrds.BackupProps{
			Retention: awscdk.Duration_Days(jsii.Number(7)),
		},
		Vpc:                     vpc,
		SubnetGroup:             subnetGrp,
		SecurityGroups:          &[]awsec2.ISecurityGroup{securityGroup},
		ServerlessV2MinCapacity: jsii.Number(0.5),
		ServerlessV2MaxCapacity: jsii.Number(2),
	})

	awscdk.NewCfnOutput(stack, jsii.String("AuroraClusterName"), &awscdk.CfnOutputProps{
		Value: databaseCluster.ClusterIdentifier(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("AuroraClusterEndpoint"), &awscdk.CfnOutputProps{
		Value: databaseCluster.ClusterEndpoint().Hostname(),
	})
}

func newLambda(stack awscdk.Stack, ecr awsecr.Repository, cfg *config.Config, vpc awsec2.IVpc, securityGroup awsec2.SecurityGroup) {
	lambdaFunc := awslambda.NewDockerImageFunction(stack, jsii.String(LambdaName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String("lambda")},
		}),
		Timeout:           awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention:      awslogs.RetentionDays_FIVE_DAYS,
		AllowPublicSubnet: jsii.Bool(true),
		Vpc:               vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetFilters: &[]awsec2.SubnetFilter{
				awsec2.SubnetFilter_ByIds(&[]*string{
					jsii.String("subnet-0901a7e554e09d234"),
					jsii.String("subnet-041de806aee128a88"),
				}),
			},
		},
		SecurityGroups: &[]awsec2.ISecurityGroup{securityGroup},
		Environment: &map[string]*string{
			"DB_DRIVER_NAME":       jsii.String(cfg.DBDriverName),
			"DB_SOURCE":            jsii.String(cfg.DBSource),
			"SESSION_SECRET_KEY":   jsii.String(cfg.SessionSecretKey),
			"JWT_SECRET_KEY":       jsii.String(cfg.JWTSecretKey),
			"GOOGLE_CLIENT_ID":     jsii.String(cfg.GoogleClientID),
			"GOOGLE_CLIENT_SECRET": jsii.String(cfg.GoogleClientSecret),
			"WEB_URL":              jsii.String(cfg.WebURL),
			"ENCRYPTION_KEY":       jsii.String(cfg.EncryptionKey),
		},
	})

	lambdaName := "falcon-cron"

	lambdaCronFunc := awslambda.NewDockerImageFunction(stack, jsii.String(lambdaName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String("lambda-cron")},
		}),
		Timeout:           awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention:      awslogs.RetentionDays_FIVE_DAYS,
		AllowPublicSubnet: jsii.Bool(true),
		Vpc:               vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetFilters: &[]awsec2.SubnetFilter{
				awsec2.SubnetFilter_ByIds(&[]*string{
					jsii.String("subnet-0901a7e554e09d234"),
					jsii.String("subnet-041de806aee128a88"),
				}),
			},
		},
		SecurityGroups: &[]awsec2.ISecurityGroup{securityGroup},
		Environment: &map[string]*string{
			"DB_DRIVER_NAME":       jsii.String(cfg.DBDriverName),
			"DB_SOURCE":            jsii.String(cfg.DBSource),
			"SESSION_SECRET_KEY":   jsii.String(cfg.SessionSecretKey),
			"JWT_SECRET_KEY":       jsii.String(cfg.JWTSecretKey),
			"GOOGLE_CLIENT_ID":     jsii.String(cfg.GoogleClientID),
			"GOOGLE_CLIENT_SECRET": jsii.String(cfg.GoogleClientSecret),
			"WEB_URL":              jsii.String(cfg.WebURL),
			"ENCRYPTION_KEY":       jsii.String(cfg.EncryptionKey),
		},
	})

	workerSQS := awssqs.NewQueue(stack, jsii.String("falcon-worker-sqs"), &awssqs.QueueProps{
		QueueName:         jsii.String("falcon-worker-sqs"),
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(500)),
	})

	lambdaWorkerName := "falcon-worker"
	lambdaWorkerFunc := awslambda.NewDockerImageFunction(stack, jsii.String(lambdaWorkerName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String("lambda-worker")},
		}),
		Timeout:           awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention:      awslogs.RetentionDays_FIVE_DAYS,
		AllowPublicSubnet: jsii.Bool(true),
		Vpc:               vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetFilters: &[]awsec2.SubnetFilter{
				awsec2.SubnetFilter_ByIds(&[]*string{
					jsii.String("subnet-0901a7e554e09d234"),
					jsii.String("subnet-041de806aee128a88"),
				}),
			},
		},
		SecurityGroups: &[]awsec2.ISecurityGroup{securityGroup},
		Environment: &map[string]*string{
			"DB_DRIVER_NAME":       jsii.String(cfg.DBDriverName),
			"DB_SOURCE":            jsii.String(cfg.DBSource),
			"SESSION_SECRET_KEY":   jsii.String(cfg.SessionSecretKey),
			"JWT_SECRET_KEY":       jsii.String(cfg.JWTSecretKey),
			"GOOGLE_CLIENT_ID":     jsii.String(cfg.GoogleClientID),
			"GOOGLE_CLIENT_SECRET": jsii.String(cfg.GoogleClientSecret),
			"WEB_URL":              jsii.String(cfg.WebURL),
			"ENCRYPTION_KEY":       jsii.String(cfg.EncryptionKey),
		},
	})

	awscdk.NewCfnOutput(stack, jsii.String("lambdaWorkerName"), &awscdk.CfnOutputProps{
		Value: lambdaWorkerFunc.FunctionName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("sqsQueueName"), &awscdk.CfnOutputProps{
		Value: workerSQS.QueueName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("sqsQueueURL"), &awscdk.CfnOutputProps{
		Value: workerSQS.QueueUrl(),
	})

	// After create sqs and lambda func then add event source.
	eventSource := awslambdaeventsources.NewSqsEventSource(workerSQS, &awslambdaeventsources.SqsEventSourceProps{})
	lambdaWorkerFunc.AddEventSource(eventSource)

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
			//	AWSLambdaVPCAccessExecutionRole
		},
	})
	lambdaPolicy.AddAllResources()
	lambdaFunc.AddToRolePolicy(lambdaPolicy)
	lambdaCronFunc.AddToRolePolicy(lambdaPolicy)
	lambdaWorkerFunc.AddToRolePolicy(lambdaPolicy)

	cronRule := awsevents.NewRule(stack, jsii.String("cat-cron-rule"), &awsevents.RuleProps{
		//Schedule every 1 hour.
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Minute: jsii.String("0"),
		}),
	})
	cronRule.AddTarget(awseventstargets.NewLambdaFunction(lambdaCronFunc, nil))

	lambdaURL := lambdaFunc.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	// Add a CloudFront distribution to route between the public directory and the Lambda function URL.
	lambdaURLDomain := awscdk.Fn_Select(jsii.Number(2), awscdk.Fn_Split(jsii.String("/"), lambdaURL.Url(), nil))
	lambdaOrigin := awscloudfrontorigins.NewHttpOrigin(lambdaURLDomain, &awscloudfrontorigins.HttpOriginProps{
		ProtocolPolicy: awscloudfront.OriginProtocolPolicy_HTTPS_ONLY,
	})

	publicDomainName := "api-falcon.vultor.xyz"
	publicDomainCertificateArn := "arn:aws:acm:us-east-1:358059338173:certificate/e74a2c12-794d-4ae4-849b-977baadf9965"

	cf := awscloudfront.NewDistribution(stack, jsii.String("customerFacing"), &awscloudfront.DistributionProps{
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

	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionUrl"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaFunctionUrl"),
		Value:      lambdaURL.Url(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionName"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaFunctionName"),
		Value:      lambdaFunc.FunctionName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("cloudFrontDomainName"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("cloudFrontDomainName"),
		Value:      cf.DomainName(),
	})
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
	repo := awsecr.NewRepository(stack, jsii.String("ECRRepository"), &awsecr.RepositoryProps{
		RepositoryName:     jsii.String(ECRName),
		RemovalPolicy:      awscdk.RemovalPolicy_DESTROY,
		ImageTagMutability: awsecr.TagMutability_MUTABLE,
		ImageScanOnPush:    jsii.Bool(false),
	})

	awscdk.NewCfnOutput(stack, jsii.String("ECRRepositoryName"), &awscdk.CfnOutputProps{
		Value: repo.RepositoryName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("ECRRepositoryURI"), &awscdk.CfnOutputProps{
		Value: repo.RepositoryUri(),
	})
	return repo
}

func newUser(stack awscdk.Stack) awsiam.User {
	user := awsiam.NewUser(stack, jsii.String("FalconUser"), &awsiam.UserProps{
		UserName: jsii.String(UserName),
	})
	awscdk.NewCfnOutput(stack, jsii.String("UserName"), &awscdk.CfnOutputProps{
		Value: user.UserName(),
	})
	awscdk.NewCfnOutput(stack, jsii.String("UserArn"), &awscdk.CfnOutputProps{
		Value: user.UserArn(),
	})

	accessKey := awsiam.NewAccessKey(stack, jsii.String("FalconUserAccessKey"), &awsiam.AccessKeyProps{
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
		&[]string{"config"}[0],
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
