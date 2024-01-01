package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
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
	newLambda(stack, ecr, cfg)

	return stack
}
func newLambda(stack awscdk.Stack, ecr awsecr.Repository, cfg *config.Config) {
	lambdaFunc := awslambda.NewDockerImageFunction(stack, jsii.String(LambdaName), &awslambda.DockerImageFunctionProps{
		Code: awslambda.DockerImageCode_FromEcr(ecr, &awslambda.EcrImageCodeProps{
			TagOrDigest: jsii.String("latest"),
			Cmd:         &[]*string{jsii.String("lambda")},
		}),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(500)),
		LogRetention: awslogs.RetentionDays_FIVE_DAYS,
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

	lambdaPolicy := awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect: awsiam.Effect_ALLOW,
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("logs:PutLogEvents"),
		},
	})
	lambdaPolicy.AddAllResources()
	lambdaFunc.AddToRolePolicy(lambdaPolicy)

	lambdaURL := lambdaFunc.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})
	awscdk.NewCfnOutput(stack, jsii.String("lambdaFunctionUrl"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("lambdaFunctionUrl"),
		Value:      lambdaURL.Url(),
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
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
