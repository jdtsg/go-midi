// Code generated by "stringer -type=GM -output=gm_string.go"; DO NOT EDIT.

package midi

import "fmt"

const _GM_name = "AcousticGrandPianoBrightAcousticPianoElectricGrandPianoHonkyTonkPianoElectricPiano1ElectricPiano2HarpsichordClaviCelestaGlockenspielMusicBoxVibraphoneMarimbaXylophoneTubularBellsDulcimerDrawbarOrganPercussiveOrganRockOrganChurchOrganReedOrganAccordionHarmonicaTangoAccordionAcousticNylonGuitarAcousticSteelGuitarElectricJazzGuitarElectricCleanGuitarElectricMutedGuitarOverdrivenGuitarDistortionGuitarGuitarHarmonicsAcousticBassElectricFingerBassElectricPickBassFretlessBassSlapBass1SlapBass2SynthBass1SynthBass2ViolinViolaCelloContrabassTremoloStringsPizzicatoStringsOrchestralHarpTimpaniStringEnsemble1StringEnsemble2SynthStrings1SynthStrings2ChoirAahsVoiceOohsSynthVoiceOrchestraHitTrumpetTromboneTubaMutedTrumpetFrenchHornBrassSectionSynthBrass1SynthBrass2SopranoSaxAltoSaxTenorSaxBaritoneSaxOboeEnglishHornBassoonClarinetPiccoloFluteRecorderPanFluteBlownBottleShakuhachiWhistleOcarinaLead1SquareLead2SawtoothLead3CalliopeLead4ChiffLead5CharangLead6VoiceLead7FifthsLead8BassLeadPad1NewAgePad2WarmPad3PolysynthPad4ChoirPad5BowedPad6MetallicPad7HaloPad8SweepFX1RainFX2SoundtrackFX3CrystalFX4AtmosphereFX5BrightnessFX6GoblinsFX7EchoesFX8SciFiSitarBanjoShamisenKotoKalimbaBagpipeFiddleShanaiTinkleBellAgogoSteelDrumsWoodblockTaikoDrumMelodicTomSynthDrumReverseCymbalGuitarFretNoiseBreathNoiseSeashoreBirdTweetTelephoneRingHelicopterApplauseGunshot"

var _GM_index = [...]uint16{0, 18, 37, 55, 69, 83, 97, 108, 113, 120, 132, 140, 150, 157, 166, 178, 186, 198, 213, 222, 233, 242, 251, 260, 274, 293, 312, 330, 349, 368, 384, 400, 415, 427, 445, 461, 473, 482, 491, 501, 511, 517, 522, 527, 537, 551, 567, 581, 588, 603, 618, 631, 644, 653, 662, 672, 684, 691, 699, 703, 715, 725, 737, 748, 759, 769, 776, 784, 795, 799, 810, 817, 825, 832, 837, 845, 853, 864, 874, 881, 888, 899, 912, 925, 935, 947, 957, 968, 981, 991, 999, 1012, 1021, 1030, 1042, 1050, 1059, 1066, 1079, 1089, 1102, 1115, 1125, 1134, 1142, 1147, 1152, 1160, 1164, 1171, 1178, 1184, 1190, 1200, 1205, 1215, 1224, 1233, 1243, 1252, 1265, 1280, 1291, 1299, 1308, 1321, 1331, 1339, 1346}

func (i GM) String() string {
	if i >= GM(len(_GM_index)-1) {
		return fmt.Sprintf("GM(%d)", i)
	}
	return _GM_name[_GM_index[i]:_GM_index[i+1]]
}
