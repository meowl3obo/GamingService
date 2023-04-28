package lolModel

type SkillDetail struct {
	CharactersSorakaSpellsSorakaRAbilitySorakaR struct {
		MScriptName string `json:"mScriptName"`
		MSpell      struct {
			MAffectsTypeFlags   int `json:"mAffectsTypeFlags"`
			MAffectsStatusFlags int `json:"mAffectsStatusFlags"`
			MRequiredUnitTags   struct {
				MTagList string `json:"mTagList"`
				Type     string `json:"__type"`
			} `json:"mRequiredUnitTags"`
			MSpellTags  []string `json:"mSpellTags"`
			MDataValues []struct {
				MName   string `json:"mName"`
				MValues []int  `json:"mValues"`
				Type    string `json:"__type"`
			} `json:"mDataValues"`
			MSpellCalculations struct {
				HealingCalc struct {
					MFormulaParts []struct {
						MDataValue   string  `json:"mDataValue,omitempty"`
						Type         string  `json:"__type"`
						MCoefficient float64 `json:"mCoefficient,omitempty"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"HealingCalc"`
				AmpedHealing struct {
					MMultiplier struct {
						MSubparts []struct {
							MNumber    int    `json:"mNumber,omitempty"`
							Type       string `json:"__type"`
							MDataValue string `json:"mDataValue,omitempty"`
						} `json:"mSubparts"`
						Type string `json:"__type"`
					} `json:"mMultiplier"`
					MModifiedGameCalculation string `json:"mModifiedGameCalculation"`
					Type                     string `json:"__type"`
				} `json:"AmpedHealing"`
			} `json:"mSpellCalculations"`
			MCoefficient              float64  `json:"mCoefficient"`
			MAnimationName            string   `json:"mAnimationName"`
			MImgIconName              []string `json:"mImgIconName"`
			MCastTime                 float64  `json:"mCastTime"`
			CooldownTime              []int    `json:"cooldownTime"`
			MCantCancelWhileWindingUp bool     `json:"mCantCancelWhileWindingUp"`
			UseAnimatorFramerate      bool     `json:"useAnimatorFramerate"`
			CastRange                 []int    `json:"castRange"`
			CastRadius                []int    `json:"castRadius"`
			CastConeDistance          int      `json:"castConeDistance"`
			CastFrame                 float64  `json:"castFrame"`
			MissileSpeed              int      `json:"missileSpeed"`
			Mana                      []int    `json:"mana"`
			SelectionPriority         int      `json:"selectionPriority"`
			MTargetingTypeData        struct {
				Type string `json:"__type"`
			} `json:"mTargetingTypeData"`
			MClientData struct {
				MTooltipData struct {
					MObjectName string `json:"mObjectName"`
					MFormat     string `json:"mFormat"`
					MLocKeys    struct {
						KeyName    string `json:"keyName"`
						KeySummary string `json:"keySummary"`
						KeyTooltip string `json:"keyTooltip"`
					} `json:"mLocKeys"`
					MLists struct {
						LevelUp struct {
							LevelCount int `json:"levelCount"`
							Elements   []struct {
								Type         string `json:"type"`
								TypeIndex    int    `json:"typeIndex,omitempty"`
								NameOverride string `json:"nameOverride,omitempty"`
								Type0        string `json:"__type"`
							} `json:"elements"`
							Type string `json:"__type"`
						} `json:"LevelUp"`
					} `json:"mLists"`
					Type string `json:"__type"`
				} `json:"mTooltipData"`
				MTargeterDefinitions []struct {
					UseCasterBoundingBox bool   `json:"useCasterBoundingBox"`
					Type                 string `json:"__type"`
				} `json:"mTargeterDefinitions"`
				Type string `json:"__type"`
			} `json:"mClientData"`
			Type string `json:"__type"`
		} `json:"mSpell"`
		Type string `json:"__type"`
	} `json:"Characters/Soraka/Spells/SorakaRAbility/SorakaR"`
	CharactersSorakaSpellsSorakaWAbilitySorakaW struct {
		MScriptName string `json:"mScriptName"`
		MSpell      struct {
			MAffectsTypeFlags int `json:"mAffectsTypeFlags"`
			MRequiredUnitTags struct {
				MTagList string `json:"mTagList"`
				Type     string `json:"__type"`
			} `json:"mRequiredUnitTags"`
			MSpellTags  []string `json:"mSpellTags"`
			MDataValues []struct {
				MName   string `json:"mName"`
				MValues []int  `json:"mValues"`
				Type    string `json:"__type"`
			} `json:"mDataValues"`
			MSpellCalculations struct {
				TotalHeal struct {
					MFormulaParts []struct {
						MDataValue   string  `json:"mDataValue,omitempty"`
						Type         string  `json:"__type"`
						MCoefficient float64 `json:"mCoefficient,omitempty"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"TotalHeal"`
				MinimumHealth struct {
					MFormulaParts []struct {
						MStat      int    `json:"mStat"`
						MDataValue string `json:"mDataValue"`
						Type       string `json:"__type"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"MinimumHealth"`
			} `json:"mSpellCalculations"`
			MAnimationName            string   `json:"mAnimationName"`
			MImgIconName              []string `json:"mImgIconName"`
			MCastTime                 float64  `json:"mCastTime"`
			CooldownTime              []int    `json:"cooldownTime"`
			DelayTotalTimePercent     float64  `json:"delayTotalTimePercent"`
			MCantCancelWhileWindingUp bool     `json:"mCantCancelWhileWindingUp"`
			UseAnimatorFramerate      bool     `json:"useAnimatorFramerate"`
			CastRange                 []int    `json:"castRange"`
			CastRadius                []int    `json:"castRadius"`
			CastConeDistance          int      `json:"castConeDistance"`
			CastFrame                 float64  `json:"castFrame"`
			MissileSpeed              int      `json:"missileSpeed"`
			MFloatVarsDecimals        []int    `json:"mFloatVarsDecimals"`
			Mana                      []int    `json:"mana"`
			SelectionPriority         int      `json:"selectionPriority"`
			MClientData               struct {
				MTooltipData struct {
					MObjectName string `json:"mObjectName"`
					MFormat     string `json:"mFormat"`
					MLocKeys    struct {
						KeyName                     string `json:"keyName"`
						KeySummary                  string `json:"keySummary"`
						KeyTooltip                  string `json:"keyTooltip"`
						KeyCost                     string `json:"keyCost"`
						KeyTooltipExtendedBelowLine string `json:"keyTooltipExtendedBelowLine"`
					} `json:"mLocKeys"`
					MLists struct {
						LevelUp struct {
							LevelCount int `json:"levelCount"`
							Elements   []struct {
								Type         string `json:"type"`
								TypeIndex    int    `json:"typeIndex,omitempty"`
								NameOverride string `json:"nameOverride,omitempty"`
								Type0        string `json:"__type"`
								Multiplier   int    `json:"multiplier,omitempty"`
								Style        int    `json:"Style,omitempty"`
							} `json:"elements"`
							Type string `json:"__type"`
						} `json:"LevelUp"`
					} `json:"mLists"`
					Type string `json:"__type"`
				} `json:"mTooltipData"`
				MTargeterDefinitions []struct {
					Type string `json:"__type"`
				} `json:"mTargeterDefinitions"`
				Type string `json:"__type"`
			} `json:"mClientData"`
			Type string `json:"__type"`
		} `json:"mSpell"`
		Type string `json:"__type"`
	} `json:"Characters/Soraka/Spells/SorakaWAbility/SorakaW"`
	CharactersSorakaSpellsSorakaEAbilitySorakaE struct {
		MScriptName string `json:"mScriptName"`
		MSpell      struct {
			MAffectsTypeFlags int      `json:"mAffectsTypeFlags"`
			MSpellTags        []string `json:"mSpellTags"`
			MDataValues       []struct {
				MName   string `json:"mName"`
				MValues []int  `json:"mValues"`
				Type    string `json:"__type"`
			} `json:"mDataValues"`
			MSpellCalculations struct {
				TotalDamage struct {
					MFormulaParts []struct {
						MDataValue   string  `json:"mDataValue,omitempty"`
						Type         string  `json:"__type"`
						MCoefficient float64 `json:"mCoefficient,omitempty"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"TotalDamage"`
			} `json:"mSpellCalculations"`
			MAnimationName            string   `json:"mAnimationName"`
			MImgIconName              []string `json:"mImgIconName"`
			MCastTime                 float64  `json:"mCastTime"`
			CooldownTime              []int    `json:"cooldownTime"`
			MCantCancelWhileWindingUp bool     `json:"mCantCancelWhileWindingUp"`
			UseAnimatorFramerate      bool     `json:"useAnimatorFramerate"`
			CastRange                 []int    `json:"castRange"`
			CastRangeDisplayOverride  []int    `json:"castRangeDisplayOverride"`
			CastRadius                []int    `json:"castRadius"`
			CastConeDistance          int      `json:"castConeDistance"`
			CastFrame                 float64  `json:"castFrame"`
			MissileSpeed              int      `json:"missileSpeed"`
			MFloatVarsDecimals        []int    `json:"mFloatVarsDecimals"`
			Mana                      []int    `json:"mana"`
			SelectionPriority         int      `json:"selectionPriority"`
			MTargetingTypeData        struct {
				Type string `json:"__type"`
			} `json:"mTargetingTypeData"`
			Zero_0F7E5Bc bool `json:"{00f7e5bc}"`
			MClientData  struct {
				MTooltipData struct {
					MObjectName string `json:"mObjectName"`
					MFormat     string `json:"mFormat"`
					MLocKeys    struct {
						KeyName    string `json:"keyName"`
						KeySummary string `json:"keySummary"`
						KeyTooltip string `json:"keyTooltip"`
					} `json:"mLocKeys"`
					MLists struct {
						LevelUp struct {
							LevelCount int `json:"levelCount"`
							Elements   []struct {
								Type         string `json:"type"`
								TypeIndex    int    `json:"typeIndex,omitempty"`
								NameOverride string `json:"nameOverride,omitempty"`
								Type0        string `json:"__type"`
							} `json:"elements"`
							Type string `json:"__type"`
						} `json:"LevelUp"`
					} `json:"mLists"`
					Type string `json:"__type"`
				} `json:"mTooltipData"`
				MTargeterDefinitions []struct {
					UseCasterBoundingBox bool   `json:"useCasterBoundingBox,omitempty"`
					Type                 string `json:"__type"`
					CenterLocator        struct {
						BasePosition int    `json:"basePosition"`
						Type         string `json:"__type"`
					} `json:"centerLocator,omitempty"`
					TextureRadiusOverrideName string `json:"textureRadiusOverrideName,omitempty"`
				} `json:"mTargeterDefinitions"`
				Type string `json:"__type"`
			} `json:"mClientData"`
			Type string `json:"__type"`
		} `json:"mSpell"`
		Type string `json:"__type"`
	} `json:"Characters/Soraka/Spells/SorakaEAbility/SorakaE"`
	CharactersSorakaSpellsSorakaQAbilitySorakaQ struct {
		MScriptName string `json:"mScriptName"`
		MSpell      struct {
			MAffectsTypeFlags int `json:"mAffectsTypeFlags"`
			MRequiredUnitTags struct {
				MTagList string `json:"mTagList"`
				Type     string `json:"__type"`
			} `json:"mRequiredUnitTags"`
			MSpellTags  []string `json:"mSpellTags"`
			MDataValues []struct {
				MName   string `json:"mName"`
				MValues []int  `json:"mValues"`
				Type    string `json:"__type"`
			} `json:"mDataValues"`
			MSpellCalculations struct {
				TotalDamage struct {
					MFormulaParts []struct {
						MDataValue   string  `json:"mDataValue,omitempty"`
						Type         string  `json:"__type"`
						MCoefficient float64 `json:"mCoefficient,omitempty"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"TotalDamage"`
				TotalHot struct {
					MFormulaParts []struct {
						MDataValue   string  `json:"mDataValue,omitempty"`
						Type         string  `json:"__type"`
						MCoefficient float64 `json:"mCoefficient,omitempty"`
					} `json:"mFormulaParts"`
					Type string `json:"__type"`
				} `json:"TotalHot"`
			} `json:"mSpellCalculations"`
			MAnimationName                  string   `json:"mAnimationName"`
			MImgIconName                    []string `json:"mImgIconName"`
			CooldownTime                    []int    `json:"cooldownTime"`
			DelayCastOffsetPercent          float64  `json:"delayCastOffsetPercent"`
			DelayTotalTimePercent           float64  `json:"delayTotalTimePercent"`
			MCantCancelWhileWindingUp       bool     `json:"mCantCancelWhileWindingUp"`
			MProjectTargetToCastRange       bool     `json:"mProjectTargetToCastRange"`
			MSpellRevealsChampion           bool     `json:"mSpellRevealsChampion"`
			UseAnimatorFramerate            bool     `json:"useAnimatorFramerate"`
			CastRange                       []int    `json:"castRange"`
			CastRangeDisplayOverride        []int    `json:"castRangeDisplayOverride"`
			CastRadius                      []int    `json:"castRadius"`
			CastConeDistance                int      `json:"castConeDistance"`
			CastTargetAdditionalUnitsRadius int      `json:"castTargetAdditionalUnitsRadius"`
			CastFrame                       float64  `json:"castFrame"`
			MissileSpeed                    int      `json:"missileSpeed"`
			MFloatVarsDecimals              []int    `json:"mFloatVarsDecimals"`
			Mana                            []int    `json:"mana"`
			SelectionPriority               int      `json:"selectionPriority"`
			MTargetingTypeData              struct {
				Type string `json:"__type"`
			} `json:"mTargetingTypeData"`
			Zero_0F7E5Bc bool `json:"{00f7e5bc}"`
			MClientData  struct {
				MTooltipData struct {
					MObjectName string `json:"mObjectName"`
					MFormat     string `json:"mFormat"`
					MLocKeys    struct {
						KeyName    string `json:"keyName"`
						KeySummary string `json:"keySummary"`
						KeyTooltip string `json:"keyTooltip"`
					} `json:"mLocKeys"`
					MLists struct {
						LevelUp struct {
							LevelCount int `json:"levelCount"`
							Elements   []struct {
								Type         string `json:"type"`
								TypeIndex    int    `json:"typeIndex,omitempty"`
								NameOverride string `json:"nameOverride,omitempty"`
								Type0        string `json:"__type"`
								Multiplier   int    `json:"multiplier,omitempty"`
								Style        int    `json:"Style,omitempty"`
							} `json:"elements"`
							Type string `json:"__type"`
						} `json:"LevelUp"`
					} `json:"mLists"`
					Type string `json:"__type"`
				} `json:"mTooltipData"`
				MTargeterDefinitions []struct {
					UseCasterBoundingBox bool   `json:"useCasterBoundingBox,omitempty"`
					Type                 string `json:"__type"`
					CenterLocator        struct {
						BasePosition int    `json:"basePosition"`
						Type         string `json:"__type"`
					} `json:"centerLocator,omitempty"`
					TextureRadiusOverrideName string `json:"textureRadiusOverrideName,omitempty"`
				} `json:"mTargeterDefinitions"`
				Type string `json:"__type"`
			} `json:"mClientData"`
			Type string `json:"__type"`
		} `json:"mSpell"`
		MBuff struct {
			MDescription       string `json:"mDescription"`
			MFloatVarsDecimals []int  `json:"mFloatVarsDecimals"`
			Type               string `json:"__type"`
		} `json:"mBuff"`
		Type string `json:"__type"`
	} `json:"Characters/Soraka/Spells/SorakaQAbility/SorakaQ"`
}
