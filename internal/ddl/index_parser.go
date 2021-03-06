
//line index_parser.rl:1
package ddl

// Code generated by go generate; DO NOT EDIT.

import (
	"bytes"
)


//line index_parser.go:13
const index_start int = 1
const index_first_final int = 53
const index_error int = 0

const index_en_main int = 1


//line index_parser.rl:12


func ParseIndex(data string) (*Index, error) {
	index := &Index{}

	buffer := &bytes.Buffer{}
	cs, eof := 0, len(data);
	p, pe := 0, eof;
	
//line index_parser.go:31
	{
	cs = index_start
	}

//line index_parser.go:36
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 55:
		goto st_case_55
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 67:
			goto st2
		case 99:
			goto st2
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st1
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 82:
			goto st3
		case 114:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch data[p] {
		case 69:
			goto st4
		case 101:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 65:
			goto st5
		case 97:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 84:
			goto st6
		case 116:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 69:
			goto st7
		case 101:
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 32 {
			goto st8
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		switch data[p] {
		case 32:
			goto st8
		case 73:
			goto st9
		case 85:
			goto st46
		case 105:
			goto st9
		case 117:
			goto st46
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st8
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		switch data[p] {
		case 78:
			goto st10
		case 110:
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 68:
			goto st11
		case 100:
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 69:
			goto st12
		case 101:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 88:
			goto st13
		case 120:
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 32 {
			goto st14
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 32 {
			goto st14
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr16
			}
		default:
			goto tr16
		}
		goto st0
tr16:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line index_parser.go:362
		switch data[p] {
		case 32:
			goto tr17
		case 95:
			goto tr16
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr17
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr16
				}
			case data[p] >= 65:
				goto tr16
			}
		default:
			goto tr16
		}
		goto st0
tr17:
//line index_parser.rl:34

		index.setName(buffer.String())
		buffer.Reset()
	
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line index_parser.go:399
		switch data[p] {
		case 32:
			goto st16
		case 79:
			goto st17
		case 111:
			goto st17
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st16
		}
		goto st0
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		switch data[p] {
		case 78:
			goto st18
		case 110:
			goto st18
		}
		goto st0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 32 {
			goto st19
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 32 {
			goto st19
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st19
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr22
			}
		default:
			goto tr22
		}
		goto st0
tr22:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line index_parser.go:468
		switch data[p] {
		case 32:
			goto tr23
		case 95:
			goto tr22
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr23
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr22
				}
			case data[p] >= 65:
				goto tr22
			}
		default:
			goto tr22
		}
		goto st0
tr23:
//line index_parser.rl:42

		index.setTable(buffer.String())
		buffer.Reset()
	
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line index_parser.go:505
		switch data[p] {
		case 32:
			goto st21
		case 40:
			goto st22
		case 85:
			goto st38
		case 117:
			goto st38
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st21
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		switch data[p] {
		case 32:
			goto st22
		case 40:
			goto tr27
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr28
			}
		default:
			goto tr28
		}
		goto st0
tr27:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line index_parser.go:555
		if data[p] == 41 {
			goto st0
		}
		goto tr29
tr29:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line index_parser.go:571
		if data[p] == 41 {
			goto tr30
		}
		goto tr29
tr30:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line index_parser.go:587
		switch data[p] {
		case 32:
			goto tr31
		case 41:
			goto tr32
		case 44:
			goto tr33
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto tr31
		}
		goto st0
tr31:
//line index_parser.rl:30

		index.addExpression(buffer.String())
		buffer.Reset()
	
	goto st26
tr45:
//line index_parser.rl:26

		index.addColumn(buffer.String())
		buffer.Reset()
	
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line index_parser.go:619
		switch data[p] {
		case 32:
			goto st26
		case 41:
			goto st53
		case 44:
			goto st33
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr37
			}
		default:
			goto tr37
		}
		goto st0
tr32:
//line index_parser.rl:30

		index.addExpression(buffer.String())
		buffer.Reset()
	
	goto st53
tr46:
//line index_parser.rl:26

		index.addColumn(buffer.String())
		buffer.Reset()
	
	goto st53
tr49:
//line index_parser.rl:38

		index.setOpClass(buffer.String())
		buffer.Reset()
	
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line index_parser.go:667
		if data[p] == 32 {
			goto st54
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st54
		}
		goto st0
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
		switch data[p] {
		case 32:
			goto st54
		case 87:
			goto st27
		case 119:
			goto st27
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st54
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 72:
			goto st28
		case 104:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 69:
			goto st29
		case 101:
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 82:
			goto st30
		case 114:
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 69:
			goto st31
		case 101:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 32 {
			goto st32
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		goto tr43
tr43:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line index_parser.go:769
		goto tr43
tr33:
//line index_parser.rl:30

		index.addExpression(buffer.String())
		buffer.Reset()
	
	goto st33
tr47:
//line index_parser.rl:26

		index.addColumn(buffer.String())
		buffer.Reset()
	
	goto st33
tr50:
//line index_parser.rl:38

		index.setOpClass(buffer.String())
		buffer.Reset()
	
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line index_parser.go:797
		if data[p] == 32 {
			goto st33
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st33
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
tr44:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line index_parser.go:825
		switch data[p] {
		case 32:
			goto tr45
		case 41:
			goto tr46
		case 44:
			goto tr47
		case 95:
			goto tr44
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr45
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr44
				}
			case data[p] >= 65:
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
tr37:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line index_parser.go:865
		switch data[p] {
		case 32:
			goto tr48
		case 41:
			goto tr49
		case 44:
			goto tr50
		case 95:
			goto tr37
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr48
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr37
				}
			case data[p] >= 65:
				goto tr37
			}
		default:
			goto tr37
		}
		goto st0
tr48:
//line index_parser.rl:38

		index.setOpClass(buffer.String())
		buffer.Reset()
	
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line index_parser.go:906
		switch data[p] {
		case 32:
			goto st36
		case 41:
			goto st53
		case 44:
			goto st33
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st36
		}
		goto st0
tr28:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line index_parser.go:930
		switch data[p] {
		case 32:
			goto tr45
		case 40:
			goto tr29
		case 41:
			goto tr46
		case 44:
			goto tr47
		case 95:
			goto tr28
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr45
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr28
				}
			case data[p] >= 65:
				goto tr28
			}
		default:
			goto tr28
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 83:
			goto st39
		case 115:
			goto st39
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 73:
			goto st40
		case 105:
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 78:
			goto st41
		case 110:
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 71:
			goto st42
		case 103:
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 32 {
			goto st43
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 32 {
			goto st43
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st43
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr57
			}
		default:
			goto tr57
		}
		goto st0
tr57:
//line index_parser.rl:22

		buffer.WriteByte(data[p])
	
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line index_parser.go:1053
		switch data[p] {
		case 32:
			goto tr58
		case 95:
			goto tr57
		}
		switch {
		case data[p] < 48:
			if 9 <= data[p] && data[p] <= 13 {
				goto tr58
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr57
				}
			case data[p] >= 65:
				goto tr57
			}
		default:
			goto tr57
		}
		goto st0
tr58:
//line index_parser.rl:46

		index.setUsing(buffer.String())
		buffer.Reset()
	
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line index_parser.go:1090
		switch data[p] {
		case 32:
			goto st45
		case 40:
			goto st22
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st45
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		switch data[p] {
		case 78:
			goto st47
		case 110:
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		switch data[p] {
		case 73:
			goto st48
		case 105:
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 81:
			goto st49
		case 113:
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 85:
			goto st50
		case 117:
			goto st50
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		switch data[p] {
		case 69:
			goto tr64
		case 101:
			goto tr64
		}
		goto st0
tr64:
//line index_parser.rl:55
 index.Unique = true 
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line index_parser.go:1170
		if data[p] == 32 {
			goto st52
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st52
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		switch data[p] {
		case 32:
			goto st52
		case 73:
			goto st9
		case 105:
			goto st9
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st52
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 55:
//line index_parser.rl:50

		index.setWhere(buffer.String())
		buffer.Reset()
	
//line index_parser.go:1261
		}
	}

	_out: {}
	}

//line index_parser.rl:83


	if cs < index_first_final {
		return nil, &parseError{
		    cs: cs,
		    data: data,
		}
	}

	return index, nil
}
