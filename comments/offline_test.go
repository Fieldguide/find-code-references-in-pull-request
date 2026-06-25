package comments

import (
	"strings"
	"testing"

	ldapi "github.com/launchdarkly/api-client-go/v15"
	"github.com/launchdarkly/find-code-references-in-pull-request/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOfflineFlagCommentRendersKeyWithoutLink(t *testing.T) {
	cfg := &config.Config{
		Offline:       true,
		LdInstance:    "https://app.launchdarkly.com",
		LdEnvironment: "production",
	}
	flag := ldapi.FeatureFlag{Key: "ai-box-annotations"}

	out, err := githubFlagComment(flag, []string{"FeatureFlag.AiBoxAnnotations"}, true, false, cfg)
	require.NoError(t, err)

	assert.NotContains(t, out, "](https://app.launchdarkly.com", "offline comment must not contain a LaunchDarkly link")
	assert.Contains(t, out, "ai-box-annotations")
	assert.Contains(t, out, "FeatureFlag.AiBoxAnnotations")
	assert.True(t, strings.HasPrefix(strings.TrimSpace(out), "| ai-box-annotations |"), "got: %s", out)
}
