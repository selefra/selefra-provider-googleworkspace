package googleworkspace_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/selefra/selefra-provider-googleworkspace/constants"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
	drive "google.golang.org/api/drive/v3"
	gmail "google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	people "google.golang.org/api/people/v1"
)

func CalendarService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*calendar.Service, error) {
	opts, err := getSessionConfig(ctx, taskClient)
	if err != nil {
		return nil, err
	}

	svc, err := calendar.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func PeopleService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*people.Service, error) {
	opts, err := getSessionConfig(ctx, taskClient)
	if err != nil {
		return nil, err
	}

	svc, err := people.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func DriveService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*drive.Service, error) {
	opts, err := getSessionConfig(ctx, taskClient)
	if err != nil {
		return nil, err
	}

	svc, err := drive.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func GmailService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*gmail.Service, error) {
	opts, err := getSessionConfig(ctx, taskClient)
	if err != nil {
		return nil, err
	}

	svc, err := gmail.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func getSessionConfig(ctx context.Context, taskClient any) ([]option.ClientOption, error) {
	opts := []option.ClientOption{}

	var credentialContent, tokenPath string
	googleworkspaceConfig := taskClient.(*Client).Config
	if googleworkspaceConfig.Credentials != constants.Constants_0 {
		credentialContent = googleworkspaceConfig.Credentials
	} else if googleworkspaceConfig.CredentialFile != constants.Constants_1 {
		credentialContent = googleworkspaceConfig.CredentialFile
	}

	if googleworkspaceConfig.TokenPath != constants.Constants_2 {
		tokenPath = googleworkspaceConfig.TokenPath
	}

	if credentialContent != constants.Constants_3 {
		ts, err := getTokenSource(ctx, taskClient)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	if tokenPath != constants.Constants_4 {
		path, err := expandPath(tokenPath)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithCredentialsFile(path))
		return opts, nil
	}

	return nil, nil
}

func getTokenSource(ctx context.Context, taskClient any) (oauth2.TokenSource, error) {

	var impersonateUser string
	googleworkspaceConfig := taskClient.(*Client).Config

	var creds string
	if googleworkspaceConfig.Credentials != constants.Constants_5 {
		creds = googleworkspaceConfig.Credentials
	} else if googleworkspaceConfig.CredentialFile != constants.Constants_6 {
		creds = googleworkspaceConfig.CredentialFile
	}

	credentialContent, err := pathOrContents(creds)
	if err != nil {
		return nil, err
	}

	if googleworkspaceConfig.ImpersonatedUserEmail != constants.Constants_7 {
		impersonateUser = googleworkspaceConfig.ImpersonatedUserEmail
	}

	if impersonateUser == constants.Constants_8 {
		return nil, errors.New(constants.Impersonateduseremailmustbeconfigured)
	}

	config, err := google.JWTConfigFromJSON(
		[]byte(credentialContent),
		calendar.CalendarReadonlyScope,
		drive.DriveReadonlyScope,
		gmail.GmailReadonlyScope,
		people.ContactsOtherReadonlyScope,
		people.ContactsReadonlyScope,
		people.DirectoryReadonlyScope,
	)
	if err != nil {
		return nil, err
	}
	config.Subject = impersonateUser

	ts := config.TokenSource(ctx)

	return ts, nil
}

func pathOrContents(poc string) (string, error) {
	if len(poc) == 0 {
		return poc, nil
	}

	path, err := expandPath(poc)
	if err != nil {
		return path, err
	}

	if _, err := os.Stat(path); err == nil {
		contents, err := os.ReadFile(path)
		if err != nil {
			return string(contents), err
		}
		return string(contents), nil
	}

	if len(path) > 1 && (path[0] == '/' || path[0] == '\\') {
		return constants.Constants_9, fmt.Errorf(constants.Snosuchfileordir, path)
	}

	return poc, nil
}

func expandPath(filePath string) (string, error) {
	path := filePath
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, err
		}
	}
	return path, nil
}
