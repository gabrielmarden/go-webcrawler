package config

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
)

type Config struct {
	URL       string
	Keyword   string
	MaxResult int
}

//Constructor for a new Config type. However, we must pass as argument the correct number of itens inside the slice.
//The slice in the argument must be the following: [0] URL, [1] KEYWORD, [2] MAX_RESULT
func NewConfig(args []string) (Config, error) {
	if len(args) < 2 {
		return Config{}, fmt.Errorf("webcrawler: missing required parameters %v", args)
	}

	return Config{
		URL:       args[0],
		Keyword:   args[1],
		MaxResult: retrieveMaxResult(args),
	}, nil
}

type ValidateFunc func(*Config) error

//Validate the config properties using a slice of validation functions of type ValidateFunc.
//There is a list of built in validation functions that can handle the most cases, but feel free to customize.
func (c *Config) Validate(funcs ...ValidateFunc) (errs []error) {
	for _, fn := range funcs {
		err := fn(c)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

//Validate the required paremeters to the Config type.
//The default parameters considered for this validations are: URL and KEYWORD.
var ValidateRequiredParameter = func(c *Config) error {
	if c.URL == "" || c.Keyword == "" {
		return fmt.Errorf("input parameteres are invalid. url=%s, keyword=%s and maxResult=%d", c.URL, c.Keyword, c.MaxResult)
	}

	return nil
}

//Validate the URL passed as input into the Config type.
var ValidateURL = func(c *Config) error {
	if _, err := url.ParseRequestURI(c.URL); err != nil {
		return fmt.Errorf("error during url validation for address=%s", c.URL)
	}

	return nil
}

//Validate the KEYWORD passed as input into the Config type.
//Here, we consider some requirements such as, the upper and lower bound for the number of characters for the KEYWORD
//Also, the KEYWORD must be a Alphanumeric characterer.
var ValidateKeyword = func(c *Config) error {
	if c.Keyword == "" {
		return fmt.Errorf("missing keyword parameter")
	}

	if !(len(c.Keyword) >= 4 && len(c.Keyword) <= 32) {
		return fmt.Errorf("keyword length out of permited range")
	}

	regex, err := regexp.Compile("^[a-zA-Z0-9]*$")
	if err != nil {
		return fmt.Errorf("error during check alphanumeric requirement for keyword=%s", c.Keyword)
	}

	if !regex.MatchString(c.Keyword) {
		return fmt.Errorf("keyword is not alphanumeric. keyword=%s", c.Keyword)
	}

	return nil
}

func retrieveMaxResult(inputs []string) int {
	if len(inputs) == 3 {
		val, err := strconv.Atoi(inputs[2])
		if err != nil {
			fmt.Printf("error parsing max result parameter. input=%s", inputs)
			return -1
		}

		return val
	}

	return -1
}
