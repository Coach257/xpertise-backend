// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/authorize/all": {
            "get": {
                "description": "获取用户申请条目",
                "tags": [
                    "admin"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取条目成功。\", \"data\": \"model.AuthorizationRequest的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/authorize/deal": {
            "post": {
                "description": "处理用户申请认证",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户申请认证ID",
                        "name": "authreq_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Accept/Reject",
                        "name": "action",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "对应作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"已通过认证请求。\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/authorize/request": {
            "post": {
                "description": "发送请求认证",
                "tags": [
                    "admin"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "身份证号",
                        "name": "citizen_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "工作单位",
                        "name": "organization",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"申请认证成功。\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/branch/comment/create": {
            "post": {
                "description": "创建一条评论",
                "tags": [
                    "branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "评论内容",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户评论成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/branch/comment/give_a_like_or_dislike": {
            "post": {
                "description": "点赞或点踩",
                "tags": [
                    "branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "评论ID",
                        "name": "comment_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "对评论的操作选择，1为点赞，2为点踩",
                        "name": "method",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户操作成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/branch/comment/list_all_comments": {
            "post": {
                "description": "列出某条文献的全部评论",
                "tags": [
                    "branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"操作成功\", \"data\": \"某文献的所有评论\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/branch/comment/operate": {
            "post": {
                "description": "操作评论",
                "tags": [
                    "branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "评论ID",
                        "name": "comment_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "对评论的操作方法，1为置顶，2为取消置顶，3为删除",
                        "name": "method",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"操作成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/branch/graph/reference": {
            "post": {
                "description": "列出某条文献的三级参考文献",
                "tags": [
                    "branch"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"操作成功\", \"data\": \"某文献的2级参考文献\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/author": {
            "post": {
                "description": "查找作者是否存在",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查询成功\", \"data\": au}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/add_to_column": {
            "post": {
                "description": "添加某篇文章到专栏",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "专栏ID",
                        "name": "column_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献标题",
                        "name": "paper_title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"添加到专栏成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/create_column": {
            "post": {
                "description": "创建一个专栏",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "专栏名字",
                        "name": "column_name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"创建专栏成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/list_all_from_column": {
            "post": {
                "description": "获取某个专栏的所有内容",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "专栏ID",
                        "name": "column_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查找成功\", \"data\": \"专栏中的所有论文ID\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/remove": {
            "post": {
                "description": "删除某条推荐",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "h-index",
                        "name": "hindex",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/remove_from_column": {
            "post": {
                "description": "删除专栏中的某条论文",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "专栏ID",
                        "name": "column_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "论文ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/column/searchcol": {
            "post": {
                "description": "返回某个作者的一个专栏",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"返回专栏成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/issettled": {
            "post": {
                "description": "判断该作者是否入驻",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"false\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/recommend/create": {
            "post": {
                "description": "创建一条推荐",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "作者名字",
                        "name": "author_name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献名",
                        "name": "paper_title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "引用次数",
                        "name": "n_citation",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "h-index",
                        "name": "hindex",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "推荐理由",
                        "name": "reason",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"推荐成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/recommend/cs/top": {
            "get": {
                "description": "获取推荐数最多的前七篇CS文献",
                "tags": [
                    "portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查找成功\", \"data\": \"前七篇文献的ID\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/recommend/main/top": {
            "get": {
                "description": "获取推荐数最多的前七篇文献",
                "tags": [
                    "portal"
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查找成功\", \"data\": \"前七篇文献的ID\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/recommend/recommends_from_one_author": {
            "post": {
                "description": "获取作者推荐的所有内容",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "作者ID",
                        "name": "author_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查找成功\", \"data\": \"作者的所有推荐\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/portal/recommend/recommends_from_one_paper": {
            "post": {
                "description": "获取所有对某文章的推荐",
                "tags": [
                    "portal"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查找成功\", \"data\": \"文献的所有推荐\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/authorize/get": {
            "post": {
                "description": "获取用户的（所有）请求认证",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"获取申请信息成功。\", \"data\": \"请求认证的所有信息。\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/authorize/read": {
            "post": {
                "description": "已读一条请求认证",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "请求认证ID",
                        "name": "authreq_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"标记已读成功！\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/favorite/add": {
            "post": {
                "description": "添加收藏",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献描述",
                        "name": "paper_info",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"收藏成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/favorite/isfav": {
            "post": {
                "description": "判断是否已经被收藏",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"true\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/favorite/list": {
            "post": {
                "description": "获取收藏列表",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"user的所有收藏\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/favorite/remove": {
            "post": {
                "description": "移除收藏",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "收藏ID",
                        "name": "favor_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"删除成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "post": {
                "description": "查看用户个人信息",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"查看用户信息成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"登录成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/modify": {
            "post": {
                "description": "修改用户信息",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "原密码",
                        "name": "password1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "新密码",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "个人信息",
                        "name": "info",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"登录成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "注册",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码1",
                        "name": "password1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码2",
                        "name": "password2",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "个人信息",
                        "name": "info",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"用户创建成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/wish/add": {
            "post": {
                "description": "添加至清单",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献ID",
                        "name": "paper_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "文献辩题",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"已添加至清单\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/wish/list": {
            "post": {
                "description": "获取收藏列表",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"查询成功\",\"data\":\"user的清单\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/wish/remove": {
            "post": {
                "description": "移出清单",
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "心愿ID",
                        "name": "wish_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true, \"message\":\"已移出清单\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "Xpertise Scholar Golang Backend",
	Description: "Xpertise Scholar",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
