package tmpl

import (
	"testing"

	"github.com/micro-plat/lib4go/assert"
)

func TestRMFSFunc(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "a_input", result: "input"},
		{input: "_input", result: "input"},
		{input: "a_b_c_input", result: "b_c_input"},
		{input: "input", result: "input"},
	}
	for _, c := range cases {
		r := rmhd(c.input)
		assert.Equal(t, c.result, r)
	}
}
func TestPascalFunc(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "a_input", result: "AInput"},
		{input: "_input", result: "Input"},
		{input: "a_b_c_input", result: "ABCInput"},
		{input: "abc_bbbc_cccc_input", result: "AbcBbbcCcccInput"},
		{input: "input", result: "Input"},
	}
	for _, c := range cases {
		r := fPascal(c.input)
		assert.Equal(t, c.result, r, c.input)
	}
}
func TestShowWordFunc(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "a,input", result: "a"},
		{input: ",input", result: ""},
		{input: "a,b_c_input", result: "a"},
		{input: "abc,bbbc_cccc_input", result: "abc"},
		{input: "input", result: "input"},
		{input: "杨,杨", result: "杨"},
	}
	for _, c := range cases {
		r := shortWord(c.input)
		assert.Equal(t, c.result, r, c.input)
	}
}
func TestMysqlTypes(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "date", result: "datetime"},
		{input: "datetime", result: "datetime"},
		{input: "decimal", result: "decimal"},
		{input: "float", result: "float"},
		{input: "int", result: "int"},
		{input: "number(1)", result: "tinyint"},
		{input: "number(2)", result: "tinyint"},
		{input: "number(3)", result: "int"},
		{input: "number(4)", result: "int"},
		{input: "number(5)", result: "int"},
		{input: "number(6)", result: "int"},
		{input: "number(7)", result: "int"},
		{input: "number(8)", result: "int"},
		{input: "number(9)", result: "int"},
		{input: "number(10)", result: "int"},
		{input: "number(11)", result: "bigint"},
		{input: "number(12)", result: "bigint"},
		{input: "number(13)", result: "bigint"},
		{input: "number(14)", result: "bigint"},
		{input: "number(15)", result: "bigint"},
		{input: "number(16)", result: "bigint"},
		{input: "number(17)", result: "bigint"},
		{input: "number(18)", result: "bigint"},
		{input: "number(19)", result: "bigint"},
		{input: "number(20)", result: "bigint"},
		{input: "number(21)", result: "bigint"},
		{input: "number(1,2)", result: "decimal(1,2)"},
		{input: "number(5,2)", result: "decimal(5,2)"},
		{input: "number(10,2)", result: "decimal(10,2)"},
		{input: "number(20,5)", result: "decimal(20,5)"},
		{input: "number(20,1)", result: "decimal(20,1)"},
		{input: "varchar(32)", result: "varchar(32)"},
		{input: "varchar(2)", result: "varchar(2)"},
	}
	for _, c := range cases {
		r := sqlType(MYSQL)(c.input)
		assert.Equal(t, c.result, r, c.input)
	}
}
func TestCodeTypes(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "date", result: "time.Time"},
		{input: "datetime", result: "time.Time"},
		{input: "decimal", result: "types.Decimal"},
		{input: "float", result: "types.Decimal"},
		{input: "int", result: "int"},
		{input: "number(1)", result: "int"},
		{input: "number(2)", result: "int"},
		{input: "number(3)", result: "int"},
		{input: "number(4)", result: "int"},
		{input: "number(5)", result: "int"},
		{input: "number(6)", result: "int"},
		{input: "number(7)", result: "int"},
		{input: "number(8)", result: "int"},
		{input: "number(9)", result: "int"},
		{input: "number(10)", result: "int"},
		{input: "number(11)", result: "int64"},
		{input: "number(12)", result: "int64"},
		{input: "number(13)", result: "int64"},
		{input: "number(14)", result: "int64"},
		{input: "number(15)", result: "int64"},
		{input: "number(16)", result: "int64"},
		{input: "number(17)", result: "int64"},
		{input: "number(18)", result: "int64"},
		{input: "number(19)", result: "int64"},
		{input: "number(20)", result: "int64"},
		{input: "number(21)", result: "int64"},
		{input: "number(1,2)", result: "types.Decimal"},
		{input: "number(5,2)", result: "types.Decimal"},
		{input: "number(10,2)", result: "types.Decimal"},
		{input: "number(20,5)", result: "types.Decimal"},
		{input: "number(20,1)", result: "types.Decimal"},
		{input: "varchar(32)", result: "string"},
		{input: "varchar(2)", result: "string"},
	}
	for _, c := range cases {
		r := codeType(c.input)
		assert.Equal(t, c.result, r, c.input)
	}
}

func TestDef(t *testing.T) {
	cases := []struct {
		input  string
		result string
	}{
		{input: "seq", result: ""},
		{input: "", result: ""},
		{input: "sysdate", result: "default current_timestamp"},
		{input: "-", result: "default:'-'"},
		{input: "10", result: "default 10"},
		{input: "abc", result: "default 'abc'"},
		{input: "efb", result: "default 'efb'"},
		{input: "a1", result: "default 'a1'"},
		{input: "d1", result: "default 'd1'"},
		{input: "e1", result: "default 'e1'"},
		{input: "杨，2ae", result: "default '杨，2ae'"},
		{input: "要", result: "default '要'"},
	}
	for _, c := range cases {
		r := defValue(c.input)
		assert.Equal(t, c.result, r, c.input)
	}
}
func TestIsCon(t *testing.T) {
	cases := []struct {
		input  string
		tp     string
		result bool
	}{
		{input: "pk", tp: "pk", result: true},
		{input: "pk,seq", tp: "pk", result: true},
		{input: "seq,pk", tp: "pk", result: true},
		{input: "seq,pk,di", tp: "pk", result: true},
		{input: "seq", tp: "seq", result: true},
		{input: "pk,seq,di", tp: "seq", result: true},
		{input: "seq,pk", tp: "seq", result: true},
		{input: "di", tp: "di", result: true},
		{input: "pk,di", tp: "di", result: true},
		{input: "di,pk", tp: "di", result: true},
		{input: "seq,di,pk", tp: "di", result: true},
		{input: "dn", tp: "dn", result: true},
		{input: "pk,dn", tp: "dn", result: true},
		{input: "dn,pk", tp: "dn", result: true},
		{input: "seq,dn,pk", tp: "dn", result: true},
		{input: "sl(a_b_c)", tp: "sl", result: true},
		{input: "pk,sl(a_b_c)", tp: "sl", result: true},
		{input: "sl(a_b_c),pk", tp: "sl", result: true},
		{input: "seq,sl(a_b_c),pk", tp: "sl", result: true},
		{input: "idx(a_b_c)", tp: "idx", result: true},
		{input: "pk,idx(a_b_c)", tp: "idx", result: true},
		{input: "idx(a_b_c),pk", tp: "idx", result: true},
		{input: "idx(a_b_c,1),pk", tp: "idx", result: true},
		{input: "idx(a_b_c,2),pk", tp: "idx", result: true},
		{input: "seq,idx(a_b_c),pk", tp: "idx", result: true},
		{input: "lcruq,idx(idx_dictionary_info_type,1)", tp: "idx", result: true},
	}
	for _, c := range cases {
		r := isCons(c.input, c.tp)
		assert.Equal(t, c.result, r, c.input)
	}
}

func TestGetIndex(t *testing.T) {
	cases := []struct {
		input  string
		index  int
		name   string
		result bool
	}{
		{input: "idx(a_b_c)", result: true, index: 0, name: "a_b_c"},
		{input: "pk,idx(a_b_c)", result: true, index: 0, name: "a_b_c"},
		{input: "idx(a_b_c),pk", result: true, index: 0, name: "a_b_c"},
		{input: "idx(a_b_c,1),pk", result: true, index: 1, name: "a_b_c"},
		{input: "idx(a_b_c,2),pk", result: true, index: 2, name: "a_b_c"},
		{input: "seq,idx(a_b_c),pk", result: true, index: 0, name: "a_b_c"},
		{input: "LCRUQ,IDX(IDX_DICTIONARY_INFO_TYPE,3)", result: true, index: 3, name: "idx_dictionary_info_type"},
	}
	for _, c := range cases {
		ok, name, index := getIndex(c.input)
		assert.Equal(t, c.result, ok, c.input)
		assert.Equal(t, c.name, name, c.input)
		assert.Equal(t, index, c.index, c.input)
	}
}
