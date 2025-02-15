package middlewares

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// ValidateRequestBody returns a Fiber middleware that parses and validates the request body.
func ValidateRequestBody(schemaFactory func() interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Create a new instance of the request struct.
		obj := schemaFactory()

		// Parse the request body into the new object.
		if err := c.BodyParser(obj); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		}

		// Validate the parsed object using the validator library.
		if err := validate.Struct(obj); err != nil {
			var errorMessages []map[string]string
			// Get the underlying type of obj (it should be a pointer to a struct).
			objValue := reflect.ValueOf(obj)
			if objValue.Kind() == reflect.Ptr {
				objValue = objValue.Elem()
			}
			objType := objValue.Type()

			// Loop over validation errors.
			if errs, ok := err.(validator.ValidationErrors); ok {
				for _, e := range errs {
					// Default field name from the struct field.
					fieldName := e.Field()

					// Try to get the json tag from the struct field.
					if field, found := objType.FieldByName(e.Field()); found {
						jsonTag := field.Tag.Get("json")
						if jsonTag != "" {
							parts := strings.Split(jsonTag, ",")
							if parts[0] != "" {
								fieldName = parts[0]
							}
						} else {
							fieldName = strings.ToLower(e.Field())
						}
					}

					// Create a custom message based on the validation tag.
					var msg string
					switch e.Tag() {
					case "required":
						msg = fieldName + " is required"
					case "min":
						msg = fieldName + " must be at least " + e.Param() + " characters long"
					case "max":
						msg = fieldName + " cannot be longer than " + e.Param() + " characters"
					case "email":
						msg = fieldName + " must be a valid email address"
					default:
						msg = fieldName + " is invalid"
					}

					errorMessages = append(errorMessages, map[string]string{
						fieldName: msg,
					})
				}
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid data",
				"errors":  errorMessages,
			})
		}

		// Store the validated object in the context for later use.
		c.Locals("validatedBody", obj)
		return c.Next()
	}
}
