package gocouchDB_test

import (
    "testing"
    "github.com/GuoJing/gocouchDB"
)

func TestMemCacheCouldSet(t *testing.T) {
    mc := gocouchDB.NewMemCache()

    int_value := 100
    bool_value := false
    string_value := "test"

    mc.Set("int_value", int_value)
    mc.Set("bool_value", bool_value)
    mc.Set("string_value", string_value)

    in_memd_int_value := mc.Get("int_value", 0)
    value := in_memd_int_value.(int)

    if value != 100{
        t.Error("value not equal to 100")
    }

    in_memd_bool_value := mc.Get("bool_value", true)
    booled_value := in_memd_bool_value.(bool)

    if booled_value != false{
        t.Error("value not equal to false")
    }

    in_memd_string_value := mc.Get("string_value", "")
    stringed_value := in_memd_string_value.(string)

    if stringed_value != "test"{
        t.Error("value not equal to test")
    }
}

func TestMemCacheCouldDelete(t *testing.T) {
    mc := gocouchDB.NewMemCache()

    int_value := 100
    mc.Set("int_value", int_value)

    in_memd_int_value := mc.Get("int_value", 0)
    value := in_memd_int_value.(int)

    if value != 100{
        t.Error("value not equal to 100")
    }

    mc.Delete("int_value")

    in_memd_int_value = mc.Get("int_value", 200)
    value = in_memd_int_value.(int)

    if value != 200{
        t.Error("value not equal to 200")
    }
}