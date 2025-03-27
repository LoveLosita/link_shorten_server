namespace go link

struct Status {
  1: required string code
  2: required string message
}

struct GenerateLinkRequest { #生成短链请求
  1: required string long_url
  2: optional string token
}

struct GenerateLinkResponse { #生成短链响应
  1: required Status status
  2: optional string short_url
}

struct DeleteLinkRequest { #删除短链请求
  1: required string short_url
  2: optional string token
}

struct DeleteLinkResponse { #删除短链响应
  1: required Status status
}

struct ChangeLinkRequest { #修改短链请求
  1: required string short_url
  2: required string new_long_url
  3: optional string token
}

struct ChangeLinkResponse { #修改短链响应
  1: required Status status
}

struct SeeLinkRankingRequest { #查看用户短链排名请求
}

struct SeeLinkRankingResponseData { #查看用户短链排名响应数据
  1: required i32 user_id
  2: required i32 rank
  3: required i32 short_url_count
}

struct SeeLinkRankingResponse { #查看用户短链排名响应
  1: required Status status
  2: optional list<SeeLinkRankingResponseData> short_url_list
}

struct SeeUserLinkRequest { #查看用户短链请求
  1: required i32 user_id
}

struct SeeUserLinkResponseData { #查看用户短链响应数据
  1: required string short_url
  2: required string long_url
}

struct SeeUserLinkResponse { #查看用户短链响应
  1: required Status status
  2: optional list<SeeUserLinkResponseData> short_url_list
}

struct LinkRedirectRequest { #短链重定向请求
  1: required string short_url
}

struct LinkRedirectResponse { #短链重定向响应
    1: required Status status
    2: optional string long_url
}

service LinkService {
  GenerateLinkResponse generate_link(1: GenerateLinkRequest req) #生成短链
  DeleteLinkResponse delete_link(1: DeleteLinkRequest req) #删除短链
  ChangeLinkResponse change_link(1: ChangeLinkRequest req) #修改短链
  SeeLinkRankingResponse see_link_ranking(1: SeeLinkRankingRequest req) #查看用户短链排名
  SeeUserLinkResponse see_user_link(1: SeeUserLinkRequest req) #查看用户短链
  LinkRedirectResponse link_redirect(1: LinkRedirectRequest req) #短链重定向
}