package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	// 做西红柿
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)
}
