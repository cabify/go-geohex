package geohex

import (
	"encoding/json"
	"log"
	"os"
)

// Test cases downloaded from http://geohex.net/testcase/v3.2.html

type (
	code2LLTestCase struct {
		code       string
		expectedLL LL
	}

	ll2codeTestCase struct {
		level        int
		ll           LL
		expectedCode string
	}

	ll2PositionTestCase struct {
		level     int
		ll        LL
		expectedX int
		expectedY int
	}

	code2PositionTestCase struct {
		code             string
		expectedPosition Position
	}

	position2hexTestCase struct {
		level        int
		x            int
		y            int
		expectedCode string
	}

	tcFieldMapping struct {
		f json.RawMessage
		v interface{}
	}
)

func loadCode2LLTestCases() []code2LLTestCase {
	var tcs []code2LLTestCase
	// http://geohex.net/testcase/hex_v3.2_test_code2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_code2HEX.json", func(raw []json.RawMessage) error {
		tc := code2LLTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.code},
			{raw[1], &tc.expectedLL.Lat},
			{raw[2], &tc.expectedLL.Lon},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

func loadLL2CodeTestCases() []ll2codeTestCase {
	var tcs []ll2codeTestCase
	// http://geohex.net/testcase/hex_v3.2_test_coord2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_coord2HEX.json", func(raw []json.RawMessage) error {
		tc := ll2codeTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.expectedCode},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

func loadLL2PositionTestCases() []ll2PositionTestCase {
	var tcs []ll2PositionTestCase
	// http://geohex.net/testcase/hex_v3.2_test_coord2XY.json
	loadTestCasesFromJson("hex_v3.2_test_coord2XY.json", func(raw []json.RawMessage) error {
		tc := ll2PositionTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.expectedX},
			{raw[4], &tc.expectedY},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

func loadCode2PositionTestCases() []code2PositionTestCase {
	var tcs []code2PositionTestCase
	// http://geohex.net/testcase/hex_v3.2_test_code2XY.json
	loadTestCasesFromJson("hex_v3.2_test_code2XY.json", func(raw []json.RawMessage) error {
		tc := code2PositionTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.code},
			{raw[1], &tc.expectedPosition.X},
			{raw[2], &tc.expectedPosition.Y},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

func loadPosition2HexTestCases() []position2hexTestCase {
	var tcs []position2hexTestCase
	// http://geohex.net/testcase/hex_v3.2_test_XY2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_XY2HEX.json", func(raw []json.RawMessage) error {
		tc := position2hexTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.x},
			{raw[2], &tc.y},
			{raw[3], &tc.expectedCode},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

func unmarshalRawFields(fields []tcFieldMapping) error {
	for _, f := range fields {
		if err := json.Unmarshal(f.f, f.v); err != nil {
			return err
		}
	}
	return nil
}

func loadTestCasesFromJson(filename string, rawUnmarshal func([]json.RawMessage) error) {
	file, err := os.Open("testdata/" + filename)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
	decoder := json.NewDecoder(file)
	rawTcs := make([][]json.RawMessage, 0)
	err = decoder.Decode(&rawTcs)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
	for _, rtc := range rawTcs {
		if err := rawUnmarshal(rtc); err != nil {
			log.Fatalf("Error: %s", err)
			return
		}
	}
}
