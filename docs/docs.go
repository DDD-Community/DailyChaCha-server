// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alert": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 네번째 - 알림설정 완료하는 API입니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "알림설정 완료 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/apple-sign-in": {
            "post": {
                "description": "Token을 받아 access token을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "애플 로그인 API",
                "parameters": [
                    {
                        "description": "애플로그인 token",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.AppleSignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/dates": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 여부를 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "온보딩 운동일정 가져오는 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListExercisedatesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 세번째 - 시간을 생성하는 API입니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "시간정하기 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "습관 일정",
                        "name": "ExerciseDate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateExerciseDateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 두번째 - 날짜를 생성하는 API입니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "날짜정하기 생성 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "요일",
                        "name": "Weekday",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateExerciseDateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/goals": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "결심하기에서 사용할 목록들을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "결심하기 목록 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/handler.goal"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 첫번째 - 결심을 생성하는 API입니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "결심하기 생성 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "결심",
                        "name": "Goal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.goal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/level": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "홈에서 사용될 레벨, 목표를 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "유저의 레벨정보를 가져오는 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetUserLevelResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/next-exercise": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "홈에서 사용될 다음 운동정보들을 들을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "유저의 다음 운동정보 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetUserNextExerciseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/objects": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "홈에서 사용될 배경, 오브젝트 목록들을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "유저의 오브젝트 목록 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListUserObjectsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/progress": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 진행상황을 반환합니다. 결심하기가 완료됐다면 'date', 날짜정하기를 완료했다면 'time', 시간정하기를 완료했다면 'alert'을 보냅니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "온보딩 진행상황 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetOnboardingProgressResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "description": "email, password를 받아 access token을 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "로그인 API",
                "parameters": [
                    {
                        "description": "유저 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.SignInUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "description": "email, password를 받아 가입합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "회원가입 API",
                "parameters": [
                    {
                        "description": "유저 정보",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.SignInUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "유저의 온보딩 여부를 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "온보딩 상태 API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.onboardingStatus"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "access token을 확인하여 해당 토큰 유저의 이메일과 user id를 반환합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "유저정보를 가져오는 API입니다.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "bearer {token}",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.AppleSignInRequest": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "handler.CreateExerciseDateRequest": {
            "type": "object",
            "properties": {
                "exercise_dates": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "handler.GetOnboardingProgressResponse": {
            "type": "object",
            "properties": {
                "progress": {
                    "type": "string"
                }
            }
        },
        "handler.GetUserLevelResponse": {
            "type": "object",
            "properties": {
                "goal": {
                    "type": "string"
                },
                "level": {
                    "type": "integer"
                }
            }
        },
        "handler.GetUserNextExerciseResponse": {
            "type": "object",
            "properties": {
                "continuity_exercise_day": {
                    "type": "integer"
                },
                "exercise_remain_time": {
                    "type": "integer"
                },
                "object_image_url": {
                    "type": "string"
                }
            }
        },
        "handler.GetUserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "handler.ListExercisedatesResponse": {
            "type": "object",
            "properties": {
                "exercise_dates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.exerciseDate"
                    }
                },
                "goal": {
                    "$ref": "#/definitions/handler.goal"
                },
                "is_all_dates_same_time": {
                    "type": "boolean"
                }
            }
        },
        "handler.ListUserObjectsResponse": {
            "type": "object",
            "properties": {
                "backgrounds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.UserBackground"
                    }
                },
                "character_image_url": {
                    "type": "string"
                },
                "has_broken_object": {
                    "type": "boolean"
                },
                "objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.UserObject"
                    }
                }
            }
        },
        "handler.SignInResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expired_at": {
                    "type": "string"
                }
            }
        },
        "handler.SignInUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateExerciseDateRequest": {
            "type": "object",
            "properties": {
                "exercise_dates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.exerciseDate"
                    }
                }
            }
        },
        "handler.UserBackground": {
            "type": "object",
            "properties": {
                "display_order": {
                    "description": "정렬 순서",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "description": "이미지 URL",
                    "type": "string"
                }
            }
        },
        "handler.UserObject": {
            "type": "object",
            "properties": {
                "display_order": {
                    "description": "정렬 순서",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "image_url": {
                    "description": "이미지 URL",
                    "type": "string"
                },
                "object_name": {
                    "description": "오브젝트 이름",
                    "type": "string"
                },
                "object_type": {
                    "description": "오브젝트 유형",
                    "type": "string"
                }
            }
        },
        "handler.exerciseDate": {
            "type": "object",
            "properties": {
                "exercise_date": {
                    "type": "integer"
                },
                "exercise_time": {
                    "type": "integer"
                }
            }
        },
        "handler.goal": {
            "type": "object",
            "properties": {
                "goal": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                }
            }
        },
        "handler.message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.onboardingStatus": {
            "type": "object",
            "properties": {
                "is_onboarding_completed": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ec2-13-209-98-22.ap-northeast-2.compute.amazonaws.com",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "DailyChaCha Sample Swagger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
