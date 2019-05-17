package ddl

// Code generated by go generate; DO NOT EDIT.

import (
	"bytes"
)

var trigger_start int = 1
var _ = trigger_start
var trigger_first_final int = 73
var _ = trigger_first_final
var trigger_error int = 0
var _ = trigger_error
var trigger_en_main int = 1
var _ = trigger_en_main

func ParseTrigger(data string) (*Trigger, error) {
	trigger := &Trigger{}

	buffer := &bytes.Buffer{}
	cs, p, pe, eof := 0, 0, len(data), -1
	{
		cs = int(trigger_start)
	}
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
		case 73:
			goto st_case_73
		case 74:
			goto st_case_74
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
		case 53:
			goto st_case_53
		case 54:
			goto st_case_54
		case 55:
			goto st_case_55
		case 56:
			goto st_case_56
		case 57:
			goto st_case_57
		case 58:
			goto st_case_58
		case 59:
			goto st_case_59
		case 60:
			goto st_case_60
		case 61:
			goto st_case_61
		case 62:
			goto st_case_62
		case 63:
			goto st_case_63
		case 64:
			goto st_case_64
		case 65:
			goto st_case_65
		case 66:
			goto st_case_66
		case 67:
			goto st_case_67
		case 68:
			goto st_case_68
		case 69:
			goto st_case_69
		case 70:
			goto st_case_70
		case 71:
			goto st_case_71
		case 72:
			goto st_case_72

		}
		goto st_out
	_st1:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof1

		}
	st_case_1:
		switch data[p] {
		case 32:
			{
				goto _st1
			}
		case 67:
			{
				goto _st2
			}
		case 99:
			{
				goto _st2
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st1
			}

		}
		{
			goto _st0
		}
	st_case_0:
	_st0:
		cs = 0
		goto _pop
	_st2:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof2

		}
	st_case_2:
		switch data[p] {
		case 82:
			{
				goto _st3
			}
		case 114:
			{
				goto _st3
			}

		}
		{
			goto _st0
		}
	_st3:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof3

		}
	st_case_3:
		switch data[p] {
		case 69:
			{
				goto _st4
			}
		case 101:
			{
				goto _st4
			}

		}
		{
			goto _st0
		}
	_st4:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof4

		}
	st_case_4:
		switch data[p] {
		case 65:
			{
				goto _st5
			}
		case 97:
			{
				goto _st5
			}

		}
		{
			goto _st0
		}
	_st5:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof5

		}
	st_case_5:
		switch data[p] {
		case 84:
			{
				goto _st6
			}
		case 116:
			{
				goto _st6
			}

		}
		{
			goto _st0
		}
	_st6:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof6

		}
	st_case_6:
		switch data[p] {
		case 69:
			{
				goto _st7
			}
		case 101:
			{
				goto _st7
			}

		}
		{
			goto _st0
		}
	_st7:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof7

		}
	st_case_7:
		if (data[p]) == 32 {
			{
				goto _st8
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st8
			}

		}
		{
			goto _st0
		}
	_st8:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof8

		}
	st_case_8:
		switch data[p] {
		case 32:
			{
				goto _st8
			}
		case 67:
			{
				goto _st9
			}
		case 84:
			{
				goto _st20
			}
		case 99:
			{
				goto _st9
			}
		case 116:
			{
				goto _st20
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st8
			}

		}
		{
			goto _st0
		}
	_st9:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof9

		}
	st_case_9:
		switch data[p] {
		case 79:
			{
				goto _st10
			}
		case 111:
			{
				goto _st10
			}

		}
		{
			goto _st0
		}
	_st10:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof10

		}
	st_case_10:
		switch data[p] {
		case 78:
			{
				goto _st11
			}
		case 110:
			{
				goto _st11
			}

		}
		{
			goto _st0
		}
	_st11:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof11

		}
	st_case_11:
		switch data[p] {
		case 83:
			{
				goto _st12
			}
		case 115:
			{
				goto _st12
			}

		}
		{
			goto _st0
		}
	_st12:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof12

		}
	st_case_12:
		switch data[p] {
		case 84:
			{
				goto _st13
			}
		case 116:
			{
				goto _st13
			}

		}
		{
			goto _st0
		}
	_st13:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof13

		}
	st_case_13:
		switch data[p] {
		case 82:
			{
				goto _st14
			}
		case 114:
			{
				goto _st14
			}

		}
		{
			goto _st0
		}
	_st14:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof14

		}
	st_case_14:
		switch data[p] {
		case 65:
			{
				goto _st15
			}
		case 97:
			{
				goto _st15
			}

		}
		{
			goto _st0
		}
	_st15:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof15

		}
	st_case_15:
		switch data[p] {
		case 73:
			{
				goto _st16
			}
		case 105:
			{
				goto _st16
			}

		}
		{
			goto _st0
		}
	_st16:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof16

		}
	st_case_16:
		switch data[p] {
		case 78:
			{
				goto _st17
			}
		case 110:
			{
				goto _st17
			}

		}
		{
			goto _st0
		}
	_st17:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof17

		}
	st_case_17:
		switch data[p] {
		case 84:
			{
				goto _ctr19
			}
		case 116:
			{
				goto _ctr19
			}

		}
		{
			goto _st0
		}
	_ctr19:
		{
			trigger.Constraint = true
		}

		goto _st18
	_st18:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof18

		}
	st_case_18:
		if (data[p]) == 32 {
			{
				goto _st19
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st19
			}

		}
		{
			goto _st0
		}
	_st19:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof19

		}
	st_case_19:
		switch data[p] {
		case 32:
			{
				goto _st19
			}
		case 84:
			{
				goto _st20
			}
		case 116:
			{
				goto _st20
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st19
			}

		}
		{
			goto _st0
		}
	_st20:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof20

		}
	st_case_20:
		switch data[p] {
		case 82:
			{
				goto _st21
			}
		case 114:
			{
				goto _st21
			}

		}
		{
			goto _st0
		}
	_st21:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof21

		}
	st_case_21:
		switch data[p] {
		case 73:
			{
				goto _st22
			}
		case 105:
			{
				goto _st22
			}

		}
		{
			goto _st0
		}
	_st22:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof22

		}
	st_case_22:
		switch data[p] {
		case 71:
			{
				goto _st23
			}
		case 103:
			{
				goto _st23
			}

		}
		{
			goto _st0
		}
	_st23:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof23

		}
	st_case_23:
		switch data[p] {
		case 71:
			{
				goto _st24
			}
		case 103:
			{
				goto _st24
			}

		}
		{
			goto _st0
		}
	_st24:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof24

		}
	st_case_24:
		switch data[p] {
		case 69:
			{
				goto _st25
			}
		case 101:
			{
				goto _st25
			}

		}
		{
			goto _st0
		}
	_st25:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof25

		}
	st_case_25:
		switch data[p] {
		case 82:
			{
				goto _st26
			}
		case 114:
			{
				goto _st26
			}

		}
		{
			goto _st0
		}
	_st26:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof26

		}
	st_case_26:
		if (data[p]) == 32 {
			{
				goto _st27
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st27
			}

		}
		{
			goto _st0
		}
	_st27:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof27

		}
	st_case_27:
		if (data[p]) == 32 {
			{
				goto _st27
			}

		}
		if (data[p]) < 65 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _st27
					}

				}
			}

		} else if (data[p]) > 90 {
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto _ctr29
					}

				}
			}

		} else {
			{
				goto _ctr29
			}

		}
		{
			goto _st0
		}
	_ctr29:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st28
	_st28:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof28

		}
	st_case_28:
		switch data[p] {
		case 32:
			{
				goto _ctr31
			}
		case 95:
			{
				goto _ctr29
			}

		}
		if (data[p]) < 48 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _ctr31
					}

				}
			}

		} else if (data[p]) > 57 {
			{
				if (data[p]) > 90 {
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto _ctr29
							}

						}
					}

				} else if (data[p]) >= 65 {
					{
						goto _ctr29
					}

				}
			}

		} else {
			{
				goto _ctr29
			}

		}
		{
			goto _st0
		}
	_ctr31:
		{
			trigger.setName(buffer.String())
			buffer.Reset()
		}

		goto _st29
	_st29:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof29

		}
	st_case_29:
		switch data[p] {
		case 32:
			{
				goto _st29
			}
		case 65:
			{
				goto _ctr33
			}
		case 66:
			{
				goto _ctr34
			}
		case 73:
			{
				goto _ctr35
			}
		case 97:
			{
				goto _ctr33
			}
		case 98:
			{
				goto _ctr34
			}
		case 105:
			{
				goto _ctr35
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st29
			}

		}
		{
			goto _st0
		}
	_ctr33:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st30
	_st30:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof30

		}
	st_case_30:
		switch data[p] {
		case 70:
			{
				goto _ctr37
			}
		case 102:
			{
				goto _ctr37
			}

		}
		{
			goto _st0
		}
	_ctr37:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st31
	_st31:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof31

		}
	st_case_31:
		switch data[p] {
		case 84:
			{
				goto _ctr39
			}
		case 116:
			{
				goto _ctr39
			}

		}
		{
			goto _st0
		}
	_ctr39:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st32
	_st32:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof32

		}
	st_case_32:
		switch data[p] {
		case 69:
			{
				goto _ctr41
			}
		case 101:
			{
				goto _ctr41
			}

		}
		{
			goto _st0
		}
	_ctr41:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st33
	_st33:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof33

		}
	st_case_33:
		switch data[p] {
		case 82:
			{
				goto _ctr43
			}
		case 114:
			{
				goto _ctr43
			}

		}
		{
			goto _st0
		}
	_ctr43:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st34
	_st34:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof34

		}
	st_case_34:
		if (data[p]) == 32 {
			{
				goto _ctr45
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _ctr45
			}

		}
		{
			goto _st0
		}
	_ctr45:
		{
			trigger.setCalled(buffer.String())
			buffer.Reset()
		}

		goto _st35
	_st35:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof35

		}
	st_case_35:
		switch data[p] {
		case 32:
			{
				goto _st35
			}
		case 68:
			{
				goto _ctr47
			}
		case 73:
			{
				goto _ctr48
			}
		case 84:
			{
				goto _ctr49
			}
		case 85:
			{
				goto _ctr50
			}
		case 100:
			{
				goto _ctr47
			}
		case 105:
			{
				goto _ctr48
			}
		case 116:
			{
				goto _ctr49
			}
		case 117:
			{
				goto _ctr50
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st35
			}

		}
		{
			goto _st0
		}
	_ctr47:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st36
	_st36:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof36

		}
	st_case_36:
		switch data[p] {
		case 69:
			{
				goto _ctr52
			}
		case 101:
			{
				goto _ctr52
			}

		}
		{
			goto _st0
		}
	_ctr52:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st37
	_st37:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof37

		}
	st_case_37:
		switch data[p] {
		case 76:
			{
				goto _ctr54
			}
		case 108:
			{
				goto _ctr54
			}

		}
		{
			goto _st0
		}
	_ctr54:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st38
	_st38:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof38

		}
	st_case_38:
		switch data[p] {
		case 69:
			{
				goto _ctr56
			}
		case 101:
			{
				goto _ctr56
			}

		}
		{
			goto _st0
		}
	_ctr56:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st39
	_st39:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof39

		}
	st_case_39:
		switch data[p] {
		case 84:
			{
				goto _ctr58
			}
		case 116:
			{
				goto _ctr58
			}

		}
		{
			goto _st0
		}
	_ctr58:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st40
	_st40:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof40

		}
	st_case_40:
		switch data[p] {
		case 69:
			{
				goto _ctr60
			}
		case 101:
			{
				goto _ctr60
			}

		}
		{
			goto _st0
		}
	_ctr60:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st41
	_st41:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof41

		}
	st_case_41:
		if (data[p]) == 32 {
			{
				goto _ctr62
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _ctr62
			}

		}
		{
			goto _st0
		}
	_ctr62:
		{
			trigger.addEvent(buffer.String())
			buffer.Reset()
		}

		goto _st42
	_st42:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof42

		}
	st_case_42:
		switch data[p] {
		case 32:
			{
				goto _st42
			}
		case 79:
			{
				goto _st43
			}
		case 111:
			{
				goto _st43
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st42
			}

		}
		{
			goto _st0
		}
	_st43:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof43

		}
	st_case_43:
		switch data[p] {
		case 78:
			{
				goto _st44
			}
		case 82:
			{
				goto _st46
			}
		case 110:
			{
				goto _st44
			}
		case 114:
			{
				goto _st46
			}

		}
		{
			goto _st0
		}
	_st44:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof44

		}
	st_case_44:
		if (data[p]) == 32 {
			{
				goto _st45
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st45
			}

		}
		{
			goto _st0
		}
	_st45:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof45

		}
	st_case_45:
		if (data[p]) == 32 {
			{
				goto _st45
			}

		}
		if (data[p]) < 65 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _st45
					}

				}
			}

		} else if (data[p]) > 90 {
			{
				if 97 <= (data[p]) && (data[p]) <= 122 {
					{
						goto _ctr68
					}

				}
			}

		} else {
			{
				goto _ctr68
			}

		}
		{
			goto _st0
		}
	_ctr68:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st73
	_ctr116:
		{
			trigger.setTable(buffer.String())
			buffer.Reset()
		}

		goto _st73
	_st73:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof73

		}
	st_case_73:
		switch data[p] {
		case 32:
			{
				goto _ctr117
			}
		case 95:
			{
				goto _ctr68
			}

		}
		if (data[p]) < 48 {
			{
				if 9 <= (data[p]) && (data[p]) <= 13 {
					{
						goto _ctr117
					}

				}
			}

		} else if (data[p]) > 57 {
			{
				if (data[p]) > 90 {
					{
						if 97 <= (data[p]) && (data[p]) <= 122 {
							{
								goto _ctr68
							}

						}
					}

				} else if (data[p]) >= 65 {
					{
						goto _ctr68
					}

				}
			}

		} else {
			{
				goto _ctr68
			}

		}
		{
			goto _st0
		}
	_ctr117:
		{
			trigger.setTable(buffer.String())
			buffer.Reset()
		}

		goto _st74
	_st74:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof74

		}
	st_case_74:
		if (data[p]) == 32 {
			{
				goto _st74
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st74
			}

		}
		{
			goto _st0
		}
	_st46:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof46

		}
	st_case_46:
		if (data[p]) == 32 {
			{
				goto _st35
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _st35
			}

		}
		{
			goto _st0
		}
	_ctr48:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st47
	_st47:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof47

		}
	st_case_47:
		switch data[p] {
		case 78:
			{
				goto _ctr70
			}
		case 110:
			{
				goto _ctr70
			}

		}
		{
			goto _st0
		}
	_ctr70:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st48
	_st48:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof48

		}
	st_case_48:
		switch data[p] {
		case 83:
			{
				goto _ctr72
			}
		case 115:
			{
				goto _ctr72
			}

		}
		{
			goto _st0
		}
	_ctr72:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st49
	_st49:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof49

		}
	st_case_49:
		switch data[p] {
		case 69:
			{
				goto _ctr74
			}
		case 101:
			{
				goto _ctr74
			}

		}
		{
			goto _st0
		}
	_ctr74:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st50
	_st50:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof50

		}
	st_case_50:
		switch data[p] {
		case 82:
			{
				goto _ctr76
			}
		case 114:
			{
				goto _ctr76
			}

		}
		{
			goto _st0
		}
	_ctr76:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st51
	_st51:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof51

		}
	st_case_51:
		switch data[p] {
		case 84:
			{
				goto _ctr60
			}
		case 116:
			{
				goto _ctr60
			}

		}
		{
			goto _st0
		}
	_ctr49:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st52
	_st52:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof52

		}
	st_case_52:
		switch data[p] {
		case 82:
			{
				goto _ctr79
			}
		case 114:
			{
				goto _ctr79
			}

		}
		{
			goto _st0
		}
	_ctr79:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st53
	_st53:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof53

		}
	st_case_53:
		switch data[p] {
		case 85:
			{
				goto _ctr81
			}
		case 117:
			{
				goto _ctr81
			}

		}
		{
			goto _st0
		}
	_ctr81:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st54
	_st54:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof54

		}
	st_case_54:
		switch data[p] {
		case 78:
			{
				goto _ctr83
			}
		case 110:
			{
				goto _ctr83
			}

		}
		{
			goto _st0
		}
	_ctr83:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st55
	_st55:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof55

		}
	st_case_55:
		switch data[p] {
		case 67:
			{
				goto _ctr85
			}
		case 99:
			{
				goto _ctr85
			}

		}
		{
			goto _st0
		}
	_ctr85:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st56
	_st56:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof56

		}
	st_case_56:
		switch data[p] {
		case 65:
			{
				goto _ctr56
			}
		case 97:
			{
				goto _ctr56
			}

		}
		{
			goto _st0
		}
	_ctr50:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st57
	_st57:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof57

		}
	st_case_57:
		switch data[p] {
		case 80:
			{
				goto _ctr88
			}
		case 112:
			{
				goto _ctr88
			}

		}
		{
			goto _st0
		}
	_ctr88:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st58
	_st58:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof58

		}
	st_case_58:
		switch data[p] {
		case 68:
			{
				goto _ctr85
			}
		case 100:
			{
				goto _ctr85
			}

		}
		{
			goto _st0
		}
	_ctr34:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st59
	_st59:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof59

		}
	st_case_59:
		switch data[p] {
		case 69:
			{
				goto _ctr91
			}
		case 101:
			{
				goto _ctr91
			}

		}
		{
			goto _st0
		}
	_ctr91:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st60
	_st60:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof60

		}
	st_case_60:
		switch data[p] {
		case 70:
			{
				goto _ctr93
			}
		case 102:
			{
				goto _ctr93
			}

		}
		{
			goto _st0
		}
	_ctr93:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st61
	_st61:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof61

		}
	st_case_61:
		switch data[p] {
		case 79:
			{
				goto _ctr95
			}
		case 111:
			{
				goto _ctr95
			}

		}
		{
			goto _st0
		}
	_ctr95:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st62
	_st62:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof62

		}
	st_case_62:
		switch data[p] {
		case 82:
			{
				goto _ctr97
			}
		case 114:
			{
				goto _ctr97
			}

		}
		{
			goto _st0
		}
	_ctr97:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st63
	_st63:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof63

		}
	st_case_63:
		switch data[p] {
		case 69:
			{
				goto _ctr43
			}
		case 101:
			{
				goto _ctr43
			}

		}
		{
			goto _st0
		}
	_ctr35:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st64
	_st64:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof64

		}
	st_case_64:
		switch data[p] {
		case 78:
			{
				goto _ctr100
			}
		case 110:
			{
				goto _ctr100
			}

		}
		{
			goto _st0
		}
	_ctr100:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st65
	_st65:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof65

		}
	st_case_65:
		switch data[p] {
		case 83:
			{
				goto _ctr102
			}
		case 115:
			{
				goto _ctr102
			}

		}
		{
			goto _st0
		}
	_ctr102:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st66
	_st66:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof66

		}
	st_case_66:
		switch data[p] {
		case 84:
			{
				goto _ctr104
			}
		case 116:
			{
				goto _ctr104
			}

		}
		{
			goto _st0
		}
	_ctr104:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st67
	_st67:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof67

		}
	st_case_67:
		switch data[p] {
		case 69:
			{
				goto _ctr106
			}
		case 101:
			{
				goto _ctr106
			}

		}
		{
			goto _st0
		}
	_ctr106:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st68
	_st68:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof68

		}
	st_case_68:
		switch data[p] {
		case 65:
			{
				goto _ctr108
			}
		case 97:
			{
				goto _ctr108
			}

		}
		{
			goto _st0
		}
	_ctr108:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st69
	_st69:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof69

		}
	st_case_69:
		switch data[p] {
		case 68:
			{
				goto _ctr110
			}
		case 100:
			{
				goto _ctr110
			}

		}
		{
			goto _st0
		}
	_ctr110:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st70
	_st70:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof70

		}
	st_case_70:
		if (data[p]) == 32 {
			{
				goto _ctr112
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _ctr112
			}

		}
		{
			goto _st0
		}
	_ctr112:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st71
	_st71:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof71

		}
	st_case_71:
		switch data[p] {
		case 32:
			{
				goto _ctr112
			}
		case 79:
			{
				goto _ctr114
			}
		case 111:
			{
				goto _ctr114
			}

		}
		if 9 <= (data[p]) && (data[p]) <= 13 {
			{
				goto _ctr112
			}

		}
		{
			goto _st0
		}
	_ctr114:
		{
			buffer.WriteByte((data[p]))
		}

		goto _st72
	_st72:
		if p == eof {
			{
				if cs >= 73 {
					goto _out

				} else {
					goto _pop

				}
			}

		}
		p += 1
		if p == pe {
			goto _test_eof72

		}
	st_case_72:
		switch data[p] {
		case 70:
			{
				goto _ctr43
			}
		case 102:
			{
				goto _ctr43
			}

		}
		{
			goto _st0
		}
	st_out:
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof
	_test_eof15:
		cs = 15
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof23:
		cs = 23
		goto _test_eof
	_test_eof24:
		cs = 24
		goto _test_eof
	_test_eof25:
		cs = 25
		goto _test_eof
	_test_eof26:
		cs = 26
		goto _test_eof
	_test_eof27:
		cs = 27
		goto _test_eof
	_test_eof28:
		cs = 28
		goto _test_eof
	_test_eof29:
		cs = 29
		goto _test_eof
	_test_eof30:
		cs = 30
		goto _test_eof
	_test_eof31:
		cs = 31
		goto _test_eof
	_test_eof32:
		cs = 32
		goto _test_eof
	_test_eof33:
		cs = 33
		goto _test_eof
	_test_eof34:
		cs = 34
		goto _test_eof
	_test_eof35:
		cs = 35
		goto _test_eof
	_test_eof36:
		cs = 36
		goto _test_eof
	_test_eof37:
		cs = 37
		goto _test_eof
	_test_eof38:
		cs = 38
		goto _test_eof
	_test_eof39:
		cs = 39
		goto _test_eof
	_test_eof40:
		cs = 40
		goto _test_eof
	_test_eof41:
		cs = 41
		goto _test_eof
	_test_eof42:
		cs = 42
		goto _test_eof
	_test_eof43:
		cs = 43
		goto _test_eof
	_test_eof44:
		cs = 44
		goto _test_eof
	_test_eof45:
		cs = 45
		goto _test_eof
	_test_eof73:
		cs = 73
		goto _test_eof
	_test_eof74:
		cs = 74
		goto _test_eof
	_test_eof46:
		cs = 46
		goto _test_eof
	_test_eof47:
		cs = 47
		goto _test_eof
	_test_eof48:
		cs = 48
		goto _test_eof
	_test_eof49:
		cs = 49
		goto _test_eof
	_test_eof50:
		cs = 50
		goto _test_eof
	_test_eof51:
		cs = 51
		goto _test_eof
	_test_eof52:
		cs = 52
		goto _test_eof
	_test_eof53:
		cs = 53
		goto _test_eof
	_test_eof54:
		cs = 54
		goto _test_eof
	_test_eof55:
		cs = 55
		goto _test_eof
	_test_eof56:
		cs = 56
		goto _test_eof
	_test_eof57:
		cs = 57
		goto _test_eof
	_test_eof58:
		cs = 58
		goto _test_eof
	_test_eof59:
		cs = 59
		goto _test_eof
	_test_eof60:
		cs = 60
		goto _test_eof
	_test_eof61:
		cs = 61
		goto _test_eof
	_test_eof62:
		cs = 62
		goto _test_eof
	_test_eof63:
		cs = 63
		goto _test_eof
	_test_eof64:
		cs = 64
		goto _test_eof
	_test_eof65:
		cs = 65
		goto _test_eof
	_test_eof66:
		cs = 66
		goto _test_eof
	_test_eof67:
		cs = 67
		goto _test_eof
	_test_eof68:
		cs = 68
		goto _test_eof
	_test_eof69:
		cs = 69
		goto _test_eof
	_test_eof70:
		cs = 70
		goto _test_eof
	_test_eof71:
		cs = 71
		goto _test_eof
	_test_eof72:
		cs = 72
		goto _test_eof

	_test_eof:
		{
		}
		if p == eof {
			{
				switch cs {
				case 1:
					{
						break
					}
				case 0:
					{
						break
					}
				case 2:
					{
						break
					}
				case 3:
					{
						break
					}
				case 4:
					{
						break
					}
				case 5:
					{
						break
					}
				case 6:
					{
						break
					}
				case 7:
					{
						break
					}
				case 8:
					{
						break
					}
				case 9:
					{
						break
					}
				case 10:
					{
						break
					}
				case 11:
					{
						break
					}
				case 12:
					{
						break
					}
				case 13:
					{
						break
					}
				case 14:
					{
						break
					}
				case 15:
					{
						break
					}
				case 16:
					{
						break
					}
				case 17:
					{
						break
					}
				case 18:
					{
						break
					}
				case 19:
					{
						break
					}
				case 20:
					{
						break
					}
				case 21:
					{
						break
					}
				case 22:
					{
						break
					}
				case 23:
					{
						break
					}
				case 24:
					{
						break
					}
				case 25:
					{
						break
					}
				case 26:
					{
						break
					}
				case 27:
					{
						break
					}
				case 28:
					{
						break
					}
				case 29:
					{
						break
					}
				case 30:
					{
						break
					}
				case 31:
					{
						break
					}
				case 32:
					{
						break
					}
				case 33:
					{
						break
					}
				case 34:
					{
						break
					}
				case 35:
					{
						break
					}
				case 36:
					{
						break
					}
				case 37:
					{
						break
					}
				case 38:
					{
						break
					}
				case 39:
					{
						break
					}
				case 40:
					{
						break
					}
				case 41:
					{
						break
					}
				case 42:
					{
						break
					}
				case 43:
					{
						break
					}
				case 44:
					{
						break
					}
				case 45:
					{
						break
					}
				case 73:
					{
						break
					}
				case 74:
					{
						break
					}
				case 46:
					{
						break
					}
				case 47:
					{
						break
					}
				case 48:
					{
						break
					}
				case 49:
					{
						break
					}
				case 50:
					{
						break
					}
				case 51:
					{
						break
					}
				case 52:
					{
						break
					}
				case 53:
					{
						break
					}
				case 54:
					{
						break
					}
				case 55:
					{
						break
					}
				case 56:
					{
						break
					}
				case 57:
					{
						break
					}
				case 58:
					{
						break
					}
				case 59:
					{
						break
					}
				case 60:
					{
						break
					}
				case 61:
					{
						break
					}
				case 62:
					{
						break
					}
				case 63:
					{
						break
					}
				case 64:
					{
						break
					}
				case 65:
					{
						break
					}
				case 66:
					{
						break
					}
				case 67:
					{
						break
					}
				case 68:
					{
						break
					}
				case 69:
					{
						break
					}
				case 70:
					{
						break
					}
				case 71:
					{
						break
					}
				case 72:
					{
						break
					}

				}
				switch cs {

				}
				switch cs {
				case 1:
					goto _st1
				case 0:
					goto _st0
				case 2:
					goto _st2
				case 3:
					goto _st3
				case 4:
					goto _st4
				case 5:
					goto _st5
				case 6:
					goto _st6
				case 7:
					goto _st7
				case 8:
					goto _st8
				case 9:
					goto _st9
				case 10:
					goto _st10
				case 11:
					goto _st11
				case 12:
					goto _st12
				case 13:
					goto _st13
				case 14:
					goto _st14
				case 15:
					goto _st15
				case 16:
					goto _st16
				case 17:
					goto _st17
				case 18:
					goto _st18
				case 19:
					goto _st19
				case 20:
					goto _st20
				case 21:
					goto _st21
				case 22:
					goto _st22
				case 23:
					goto _st23
				case 24:
					goto _st24
				case 25:
					goto _st25
				case 26:
					goto _st26
				case 27:
					goto _st27
				case 28:
					goto _st28
				case 29:
					goto _st29
				case 30:
					goto _st30
				case 31:
					goto _st31
				case 32:
					goto _st32
				case 33:
					goto _st33
				case 34:
					goto _st34
				case 35:
					goto _st35
				case 36:
					goto _st36
				case 37:
					goto _st37
				case 38:
					goto _st38
				case 39:
					goto _st39
				case 40:
					goto _st40
				case 41:
					goto _st41
				case 42:
					goto _st42
				case 43:
					goto _st43
				case 44:
					goto _st44
				case 45:
					goto _st45
				case 73:
					goto _ctr116
				case 74:
					goto _st74
				case 46:
					goto _st46
				case 47:
					goto _st47
				case 48:
					goto _st48
				case 49:
					goto _st49
				case 50:
					goto _st50
				case 51:
					goto _st51
				case 52:
					goto _st52
				case 53:
					goto _st53
				case 54:
					goto _st54
				case 55:
					goto _st55
				case 56:
					goto _st56
				case 57:
					goto _st57
				case 58:
					goto _st58
				case 59:
					goto _st59
				case 60:
					goto _st60
				case 61:
					goto _st61
				case 62:
					goto _st62
				case 63:
					goto _st63
				case 64:
					goto _st64
				case 65:
					goto _st65
				case 66:
					goto _st66
				case 67:
					goto _st67
				case 68:
					goto _st68
				case 69:
					goto _st69
				case 70:
					goto _st70
				case 71:
					goto _st71
				case 72:
					goto _st72
				}
			}

		}
		if cs >= 73 {
			goto _out
		}
	_pop:
		{
		}
	_out:
		{
		}
	}
	if cs < trigger_first_final {
		return nil, &parseError{
			cs:   cs,
			data: data,
		}
	}

	return trigger, nil
}
