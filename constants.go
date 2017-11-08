//go:generate stringer -type=Note -output=note_string.go
//go:generate stringer -type=GM -output=gm_string.go
package midi

const (
	NoteOff           = 0x80
	NoteOn            = 0x90
	NoteAfterTouch    = 0xa0
	Controller        = 0xb0
	ProgramChange     = 0xc0
	ChannelAfterTouch = 0xd0
	PitchBend         = 0xe0
	SystemExclusive   = 0xf0
	EndOfNormalSysEx  = 0xf7
)

const (
	Text                = 0x01
	CopyrightNotice     = 0x02
	SequenceOrTrackName = 0x03
	InstrumentName      = 0x04
	Lyrics              = 0x05
	Marker              = 0x06
	CuePoint            = 0x07
	MIDIPortPrefix      = 0x20
	MIDIChannelPrefix   = 0x21
	SetTempo            = 0x51
	SMPTEOffset         = 0x54
	TimeSignature       = 0x58
	KeySignature        = 0x59
	SequencerSpecific   = 0x7f
	EndOfTrack          = 0x2f
)

// Note
type Note byte

const (
	Cminus2  Note = 0x00
	Dbminus2 Note = 0x01
	Dminus2  Note = 0x02
	Ebminus2 Note = 0x03
	Eminus2  Note = 0x04
	Fminus2  Note = 0x05
	Gbminus2 Note = 0x06
	Gminus2  Note = 0x07
	Abminus2 Note = 0x08
	Aminus2  Note = 0x09
	Bbminus2 Note = 0x0a
	Bminus2  Note = 0x0b
	Cminus1  Note = 0x0c
	Dbminus1 Note = 0x0d
	Dminus1  Note = 0x0e
	Ebminus1 Note = 0x0f
	Eminus1  Note = 0x10
	Fminus1  Note = 0x11
	Gbminus1 Note = 0x12
	Gminus1  Note = 0x13
	Abminus1 Note = 0x14
	Aminus1  Note = 0x15
	Bbminus1 Note = 0x16
	Bminus1  Note = 0x17
	C0       Note = 0x18
	Db0      Note = 0x19
	D0       Note = 0x1a
	Eb0      Note = 0x1b
	E0       Note = 0x1c
	F0       Note = 0x1d
	Gb0      Note = 0x1e
	G0       Note = 0x1f
	Ab0      Note = 0x20
	A0       Note = 0x21
	Bb0      Note = 0x22
	B0       Note = 0x23
	C1       Note = 0x24
	Db1      Note = 0x25
	D1       Note = 0x26
	Eb1      Note = 0x27
	E1       Note = 0x28
	F1       Note = 0x29
	Gb1      Note = 0x2a
	G1       Note = 0x2b
	Ab1      Note = 0x2c
	A1       Note = 0x2d
	Bb1      Note = 0x2e
	B1       Note = 0x2f
	C2       Note = 0x30
	Db2      Note = 0x31
	D2       Note = 0x32
	Eb2      Note = 0x33
	E2       Note = 0x34
	F2       Note = 0x35
	Gb2      Note = 0x36
	G2       Note = 0x37
	Ab2      Note = 0x38
	A2       Note = 0x39
	Bb2      Note = 0x3a
	B2       Note = 0x3b
	C3       Note = 0x3c
	Db3      Note = 0x3d
	D3       Note = 0x3e
	Eb3      Note = 0x3f
	E3       Note = 0x40
	F3       Note = 0x41
	Gb3      Note = 0x42
	G3       Note = 0x43
	Ab3      Note = 0x44
	A3       Note = 0x45
	Bb3      Note = 0x46
	B3       Note = 0x47
	C4       Note = 0x48
	Db4      Note = 0x49
	D4       Note = 0x4a
	Eb4      Note = 0x4b
	E4       Note = 0x4c
	F4       Note = 0x4d
	Gb4      Note = 0x4e
	G4       Note = 0x4f
	Ab4      Note = 0x50
	A4       Note = 0x51
	Bb4      Note = 0x52
	B4       Note = 0x53
	C5       Note = 0x54
	Db5      Note = 0x55
	D5       Note = 0x56
	Eb5      Note = 0x57
	E5       Note = 0x58
	F5       Note = 0x59
	Gb5      Note = 0x5a
	G5       Note = 0x5b
	Ab5      Note = 0x5c
	A5       Note = 0x5d
	Bb5      Note = 0x5e
	B5       Note = 0x5f
	C6       Note = 0x60
	Db6      Note = 0x61
	D6       Note = 0x62
	Eb6      Note = 0x63
	E6       Note = 0x64
	F6       Note = 0x65
	Gb6      Note = 0x66
	G6       Note = 0x67
	Ab6      Note = 0x68
	A6       Note = 0x69
	Bb6      Note = 0x6a
	B6       Note = 0x6b
	C7       Note = 0x6c
	Db7      Note = 0x6d
	D7       Note = 0x6e
	Eb7      Note = 0x6f
	E7       Note = 0x70
	F7       Note = 0x71
	Gb7      Note = 0x72
	G7       Note = 0x73
	Ab7      Note = 0x74
	A7       Note = 0x75
	Bb7      Note = 0x76
	B7       Note = 0x77
	C8       Note = 0x78
	Db8      Note = 0x79
	D8       Note = 0x7a
	Eb8      Note = 0x7b
	E8       Note = 0x7c
	F8       Note = 0x7d
	Gb8      Note = 0x7e
	G8       Note = 0x7f
)

type Rhythm uint16

const (
	AcousticBassDrum Rhythm = 0x35
	BassDrum1        Rhythm = 0x36
	SideStick        Rhythm = 0x37
	AcousticSnare    Rhythm = 0x38
	HandClap         Rhythm = 0x39
	ElectricSnare    Rhythm = 0x40
	LowFloorTom      Rhythm = 0x41
	ClosedHiHat      Rhythm = 0x42
	HighFloorTom     Rhythm = 0x43
	PedalHiHat       Rhythm = 0x44
	LowTom           Rhythm = 0x45
	OpenHiHat        Rhythm = 0x46
	LowMidTom        Rhythm = 0x47
	HiMidTom         Rhythm = 0x48
	CrashCymbal1     Rhythm = 0x49
	HighTom          Rhythm = 0x50
	RideCymbal1      Rhythm = 0x51
	ChineseCymbal    Rhythm = 0x52
	RideBell         Rhythm = 0x53
	Tambourine       Rhythm = 0x54
	SplashCymbal     Rhythm = 0x55
	Cowbell          Rhythm = 0x56
	CrashCymbal2     Rhythm = 0x57
	Vibraslap        Rhythm = 0x58
	RideCymbal2      Rhythm = 0x59
	HiBongo          Rhythm = 0x60
	LowBongo         Rhythm = 0x61
	MuteHiConga      Rhythm = 0x62
	OpenHiConga      Rhythm = 0x63
	LowConga         Rhythm = 0x64
	HighTimbale      Rhythm = 0x65
	LowTimbale       Rhythm = 0x66
	HighAgogo        Rhythm = 0x67
	LowAgogo         Rhythm = 0x68
	Cabasa           Rhythm = 0x69
	Maracas          Rhythm = 0x70
	ShortWhistle     Rhythm = 0x71
	LongWhistle      Rhythm = 0x72
	ShortGuiro       Rhythm = 0x73
	LongGuiro        Rhythm = 0x74
	Claves           Rhythm = 0x75
	HiWoodBlock      Rhythm = 0x76
	LowWoodBlock     Rhythm = 0x77
	MuteCuica        Rhythm = 0x78
	OpenCuica        Rhythm = 0x79
	MuteTriangle     Rhythm = 0x80
	OpenTriangle     Rhythm = 0x81
)

type GM uint16

const (
	AcousticGrandPiano  GM = 0x00
	BrightAcousticPiano GM = 0x01
	ElectricGrandPiano  GM = 0x02
	HonkyTonkPiano      GM = 0x03
	ElectricPiano1      GM = 0x04
	ElectricPiano2      GM = 0x05
	Harpsichord         GM = 0x06
	Clavi               GM = 0x07
	Celesta             GM = 0x08
	Glockenspiel        GM = 0x09
	MusicBox            GM = 0x0a
	Vibraphone          GM = 0x0b
	Marimba             GM = 0x0c
	Xylophone           GM = 0x0d
	TubularBells        GM = 0x0e
	Dulcimer            GM = 0x0f
	DrawbarOrgan        GM = 0x10
	PercussiveOrgan     GM = 0x11
	RockOrgan           GM = 0x12
	ChurchOrgan         GM = 0x13
	ReedOrgan           GM = 0x14
	Accordion           GM = 0x15
	Harmonica           GM = 0x16
	TangoAccordion      GM = 0x17
	AcousticNylonGuitar GM = 0x18
	AcousticSteelGuitar GM = 0x19
	ElectricJazzGuitar  GM = 0x1a
	ElectricCleanGuitar GM = 0x1b
	ElectricMutedGuitar GM = 0x1c
	OverdrivenGuitar    GM = 0x1d
	DistortionGuitar    GM = 0x1e
	GuitarHarmonics     GM = 0x1f
	AcousticBass        GM = 0x20
	ElectricFingerBass  GM = 0x21
	ElectricPickBass    GM = 0x22
	FretlessBass        GM = 0x23
	SlapBass1           GM = 0x24
	SlapBass2           GM = 0x25
	SynthBass1          GM = 0x26
	SynthBass2          GM = 0x27
	Violin              GM = 0x28
	Viola               GM = 0x29
	Cello               GM = 0x2a
	Contrabass          GM = 0x2b
	TremoloStrings      GM = 0x2c
	PizzicatoStrings    GM = 0x2d
	OrchestralHarp      GM = 0x2e
	Timpani             GM = 0x2f
	StringEnsemble1     GM = 0x30
	StringEnsemble2     GM = 0x31
	SynthStrings1       GM = 0x32
	SynthStrings2       GM = 0x33
	ChoirAahs           GM = 0x34
	VoiceOohs           GM = 0x35
	SynthVoice          GM = 0x36
	OrchestraHit        GM = 0x37
	Trumpet             GM = 0x38
	Trombone            GM = 0x39
	Tuba                GM = 0x3a
	MutedTrumpet        GM = 0x3b
	FrenchHorn          GM = 0x3c
	BrassSection        GM = 0x3d
	SynthBrass1         GM = 0x3e
	SynthBrass2         GM = 0x3f
	SopranoSax          GM = 0x40
	AltoSax             GM = 0x41
	TenorSax            GM = 0x42
	BaritoneSax         GM = 0x43
	Oboe                GM = 0x44
	EnglishHorn         GM = 0x45
	Bassoon             GM = 0x46
	Clarinet            GM = 0x47
	Piccolo             GM = 0x48
	Flute               GM = 0x49
	Recorder            GM = 0x4a
	PanFlute            GM = 0x4b
	BlownBottle         GM = 0x4c
	Shakuhachi          GM = 0x4d
	Whistle             GM = 0x4e
	Ocarina             GM = 0x4f
	Lead1Square         GM = 0x50
	Lead2Sawtooth       GM = 0x51
	Lead3Calliope       GM = 0x52
	Lead4Chiff          GM = 0x53
	Lead5Charang        GM = 0x54
	Lead6Voice          GM = 0x55
	Lead7Fifths         GM = 0x56
	Lead8BassLead       GM = 0x57
	Pad1NewAge          GM = 0x58
	Pad2Warm            GM = 0x59
	Pad3Polysynth       GM = 0x5a
	Pad4Choir           GM = 0x5b
	Pad5Bowed           GM = 0x5c
	Pad6Metallic        GM = 0x5d
	Pad7Halo            GM = 0x5e
	Pad8Sweep           GM = 0x5f
	FX1Rain             GM = 0x60
	FX2Soundtrack       GM = 0x61
	FX3Crystal          GM = 0x62
	FX4Atmosphere       GM = 0x63
	FX5Brightness       GM = 0x64
	FX6Goblins          GM = 0x65
	FX7Echoes           GM = 0x66
	FX8SciFi            GM = 0x67
	Sitar               GM = 0x68
	Banjo               GM = 0x69
	Shamisen            GM = 0x6a
	Koto                GM = 0x6b
	Kalimba             GM = 0x6c
	Bagpipe             GM = 0x6d
	Fiddle              GM = 0x6e
	Shanai              GM = 0x6f
	TinkleBell          GM = 0x70
	Agogo               GM = 0x71
	SteelDrums          GM = 0x72
	Woodblock           GM = 0x73
	TaikoDrum           GM = 0x74
	MelodicTom          GM = 0x75
	SynthDrum           GM = 0x76
	ReverseCymbal       GM = 0x77
	GuitarFretNoise     GM = 0x78
	BreathNoise         GM = 0x79
	Seashore            GM = 0x7a
	BirdTweet           GM = 0x7b
	TelephoneRing       GM = 0x7c
	Helicopter          GM = 0x7d
	Applause            GM = 0x7e
	Gunshot             GM = 0x7f
)
