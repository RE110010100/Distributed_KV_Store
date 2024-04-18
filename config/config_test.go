package config_test

import (
	"DistributedKV/config"
	"os"
	"reflect"
	"testing"
)

func createConfig(t *testing.T, contents string) config.Config {

	t.Helper()

	f, err := os.CreateTemp(os.TempDir(), "config.toml")
	if err != nil {
		t.Fatalf("Couldnt create a temp file")
	}
	defer f.Close()

	name := f.Name()
	defer os.Remove(name)

	_, err = f.WriteString(contents)
	if err != nil {
		t.Fatalf("Could not write the config file contents: %v", err)
	}

	c, err := config.ParseFile(name)
	if err != nil {
		t.Fatalf("Could not parse config : %v", name)
	}

	return c

}

func TestConfigParse(t *testing.T) {

	got := createConfig(t, `[[shard]]
	name = "Boulder"
	idx = 0
	address = "localhost:8081"`)

	want := config.Config{
		Shard: []config.Shard{
			{
				Name:    "Boulder",
				Idx:     0,
				Address: "localhost:8081",
			},
		},
	}

	if !(reflect.DeepEqual(got, want)) {
		t.Fatalf("The config does not match; got : %v, want : %v", got, want)
	}

}

func TestParseShards(t *testing.T) {

	c := createConfig(t, `
	[[shard]]
		name = "Boulder"
		idx = 0
		address = "localhost:8080"
	[[shard]]
		name = "Bangalore"
		idx = 1
		address = "localhost:8081"`)

	got, err := config.ParseShards(c.Shard, "Boulder")
	if err != nil {
		t.Fatalf("Could not parse shards %#v: %v", c.Shard, err)
	}

	want := &config.Shards{
		Count:  2,
		CurIdx: 0,
		Addrs: map[int]string{
			0: "localhost:8080",
			1: "localhost:8081",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("The shards config does match: got: %#v, want: %#v", got, want)
	}

}
