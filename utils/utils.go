package utils

import (
	"fmt"
	"os"
	"strconv"
)

// GetEnv returns an environment variable or the given default value if it isn't defined
// All environment variables are returned as string as must be parsed to the appropiate
// type
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//GetIntEnv convert to int
func GetIntEnv(key string, defaultValue int) int {
	fmt.Printf("Trying to get environment variable %v  \n", key)
	if value, ok := os.LookupEnv(key); ok {
		convertedValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("error parsing value of key %v value  %v  \n", key, value)
			fmt.Printf("using default value environment variable %v value %v \n ", key, defaultValue)
			return defaultValue
		}
		fmt.Printf("key found environment variable %v value %v  \n", key, value)
		return convertedValue
	}
	fmt.Printf("key not found using default value environment variable %v value %v \n ", key, defaultValue)
	return defaultValue
}

//GetBoolEnv convert to bool
func GetBoolEnv(key string, defaultValue bool) bool {
	fmt.Printf("Trying to get environment variable %v  \n", key)
	if value, ok := os.LookupEnv(key); ok {
		convertedValue, err := strconv.ParseBool(value)
		if err != nil {
			fmt.Printf("error parsing value of key %v value  %v  \n", key, value)
			fmt.Printf("using default value environment variable %v value %v \n ", key, defaultValue)
			return defaultValue
		}
		fmt.Printf("key found environment variable %v value %v  \n", key, value)
		return convertedValue
	}
	fmt.Printf("key not found using default value environment variable %v value %v \n ", key, defaultValue)
	return defaultValue
}
