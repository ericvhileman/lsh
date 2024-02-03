package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/latitudesh/lsh/client/ssh_keys"
	"github.com/latitudesh/lsh/internal/utils"
	"github.com/latitudesh/lsh/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSSHKeysPostProjectSSHKeyCmd returns a cmd to handle operation postProjectSshKey
func makeOperationSSHKeysPostProjectSSHKeyCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "create",
		Short: `Allow you create SSH Keys in a project. These keys can be used to access servers after deploy and reinstall actions.`,
		RunE:  runOperationSSHKeysPostProjectSSHKey,
	}

	if err := registerOperationSSHKeysPostProjectSSHKeyParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSSHKeysPostProjectSSHKey uses cmd flags to call endpoint api
func runOperationSSHKeysPostProjectSSHKey(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ssh_keys.NewPostProjectSSHKeyParams()
	if err, _ := retrieveOperationSSHKeysPostProjectSSHKeyBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}

	response, err := appCli.SSHKeys.PostProjectSSHKey(params, nil)
	if err != nil {
		utils.PrintError(err)
		return nil
	}

	if !debug {
		utils.Render(response.GetPayload())
	}
	return nil
}

// registerOperationSSHKeysPostProjectSSHKeyParamFlags registers all flags needed to fill params
func registerOperationSSHKeysPostProjectSSHKeyParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSSHKeysPostProjectSSHKeyBodyParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSSHKeysPostProjectSSHKeyBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. ")

	// add flags for body
	if err := registerModelPostProjectSSHKeyBodyFlags(0, "", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	projectIdOrSlugDescription := `Project Id or Slug (Required).`

	var projectIdOrSlugFlagName string
	if cmdPrefix == "" {
		projectIdOrSlugFlagName = "project"
	} else {
		projectIdOrSlugFlagName = fmt.Sprintf("%v.project", cmdPrefix)
	}

	var projectIdOrSlugFlagDefault string

	_ = cmd.PersistentFlags().String(projectIdOrSlugFlagName, projectIdOrSlugFlagDefault, projectIdOrSlugDescription)
	cmd.MarkPersistentFlagRequired(projectIdOrSlugFlagName)

	return nil
}

func retrieveOperationSSHKeysPostProjectSSHKeyBodyFlag(m *ssh_keys.PostProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := ssh_keys.PostProjectSSHKeyBody{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in PostProjectSSHKeyBody: %v", err), false
		}
		m.Body = bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = ssh_keys.PostProjectSSHKeyBody{}
	}
	err, added := retrieveModelPostProjectSSHKeyBodyFlags(0, &bodyValueModel, "", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationSSHKeysPostProjectSSHKeyProjectIDOrSlugFlag(m *ssh_keys.PostProjectSSHKeyParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("project") {

		var projectIdOrSlugFlagName string
		if cmdPrefix == "" {
			projectIdOrSlugFlagName = "project"
		} else {
			projectIdOrSlugFlagName = fmt.Sprintf("%v.project", cmdPrefix)
		}

		projectIdOrSlugFlagValue, err := cmd.Flags().GetString(projectIdOrSlugFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectIDOrSlug = projectIdOrSlugFlagValue

	}
	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelPostProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyBodyFlags(depth int, m *ssh_keys.PostProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostProjectSSHKeyBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%vdata", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data PostProjectSSHKeyParamsBodyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &ssh_keys.PostProjectSSHKeyParamsBodyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelPostProjectSSHKeyParamsBodyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyCreatedBodyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyCreatedBodyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyCreatedBodyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName = ""

	if err := registerModelSSHKeyDataFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyCreatedBodyFlags(depth int, m *ssh_keys.PostProjectSSHKeyCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataAdded := retrievePostProjectSSHKeyCreatedBodyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyCreatedBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyCreatedBody, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%vdata", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data models.SSHKeyData is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.SSHKeyData{}
	}

	dataFlagName = ""
	err, dataAdded := retrieveModelSSHKeyDataFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyParamsBodyDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyParamsBodyDataAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var attributesFlagName = ""

	if err := registerModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyParamsBodyDataFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributesAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%vattributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// info: complex object attributes PostProjectSSHKeyParamsBodyDataAttributes is retrieved outside this Changed() block
	}
	attributesFlagValue := m.Attributes
	if swag.IsZero(attributesFlagValue) {
		attributesFlagValue = &ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes{}
	}

	attributesFlagName = ""
	err, attributesAdded := retrieveModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth+1, attributesFlagValue, attributesFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded
	if attributesAdded {
		m.Attributes = attributesFlagValue
	}

	return nil, retAdded
}

// register flags to command
func registerModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPostProjectSSHKeyParamsBodyDataAttributesName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPostProjectSSHKeyParamsBodyDataAttributesPublicKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributesName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Name of the SSH Key`

	var nameFlagName = "name"

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerPostProjectSSHKeyParamsBodyDataAttributesPublicKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	publicKeyDescription := `SSH Public Key`

	var publicKeyFlagName = "public_key"

	var publicKeyFlagDefault string

	_ = cmd.PersistentFlags().String(publicKeyFlagName, publicKeyFlagDefault, publicKeyDescription)
	cmd.MarkPersistentFlagRequired(publicKeyFlagName)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPostProjectSSHKeyParamsBodyDataAttributesFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, publicKeyAdded := retrievePostProjectSSHKeyParamsBodyDataAttributesPublicKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || publicKeyAdded

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesNameFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var nameFlagName = "name"
	if cmd.Flags().Changed(nameFlagName) {

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePostProjectSSHKeyParamsBodyDataAttributesPublicKeyFlags(depth int, m *ssh_keys.PostProjectSSHKeyParamsBodyDataAttributes, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	var publicKeyFlagName = "public_key"
	if cmd.Flags().Changed(publicKeyFlagName) {

		publicKeyFlagValue, err := cmd.Flags().GetString(publicKeyFlagName)
		if err != nil {
			return err, false
		}
		m.PublicKey = publicKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}
