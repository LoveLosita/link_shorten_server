package main

import (
	"context"
	"link_shorten_server/link/dao"
	link "link_shorten_server/link/kitex_gen/link"
	"link_shorten_server/link/response"
	"link_shorten_server/user/kitex_gen/user"
	"link_shorten_server/utils"
)

// LinkServiceImpl implements the last service interface defined in the IDL.
type LinkServiceImpl struct{}

// GenerateLink implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) GenerateLink(ctx context.Context, req *link.GenerateLinkRequest) (resp *link.GenerateLinkResponse, err error) {
	var emptyUserStatus user.Status //此处使用user的Status类型是因为验证jwt的函数使用的返回类型是user.Status
	var emptyStatus link.Status
	var userID int
	var userStatus user.Status
	var status link.Status
	var shortCode string
	var generateLinkResp link.GenerateLinkResponse
	//1.获取验证用户的token，以获取userid（如果有）
	if req != nil && req.Token != nil {
		userID, userStatus = utils.CheckJwtToken(*req.Token)
		if userID == 0 || userStatus != emptyUserStatus {
			return &link.GenerateLinkResponse{Status: (*link.Status)(&userStatus)}, nil
		}
	} else if req == nil {
		return &link.GenerateLinkResponse{Status: &response.EmptyRequest}, nil
	}
	//2.获取url，并生成短链，存入数据库
	//2.1.先检查url是否已经存在
	result, shortCode, status := dao.IfLinkExists(req.LongUrl)
	if status != emptyStatus {
		return &link.GenerateLinkResponse{Status: &status}, nil
	}
	if result { //已经存在
		shortCode = "127.0.0.1/" + shortCode
		generateLinkResp.ShortUrl = &shortCode
		generateLinkResp.Status = &response.Ok
		return &generateLinkResp, nil
	}
	//2.1.将初始url先存入数据库
	linkID, status := dao.InsertLink1(req.LongUrl, userID)
	if status != emptyStatus {
		return &link.GenerateLinkResponse{Status: &status}, nil
	}
	//2.2.使用link的id生成短链
	shortCode, err = utils.IDToAbc(linkID)
	if err != nil {
		RetErr := response.InternalErr(err)
		return &link.GenerateLinkResponse{Status: &RetErr}, nil
	}
	//2.3.将生成的短链更新到数据库中
	status = dao.UpdateLink2(shortCode, linkID)
	if status != emptyStatus {
		return &link.GenerateLinkResponse{Status: &status}, nil
	}
	//3.如果用户登录了，在榜单上添加用户的短链数量
	status = dao.UpdateLinkRanking(userID)
	if status != emptyStatus {
		return &link.GenerateLinkResponse{Status: &status}, nil
	}
	//4.返回状态和短链
	shortCode = "127.0.0.1/" + shortCode
	generateLinkResp.ShortUrl = &shortCode
	generateLinkResp.Status = &response.Ok
	return &generateLinkResp, nil
}

// DeleteLink implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) DeleteLink(ctx context.Context, req *link.DeleteLinkRequest) (resp *link.DeleteLinkResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeLink implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) ChangeLink(ctx context.Context, req *link.ChangeLinkRequest) (resp *link.ChangeLinkResponse, err error) {
	// TODO: Your code here...
	return
}

// SeeLinkRanking implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) SeeLinkRanking(ctx context.Context, req *link.SeeLinkRankingRequest) (resp *link.SeeLinkRankingResponse, err error) {
	// TODO: Your code here...
	return
}

// SeeUserLink implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) SeeUserLink(ctx context.Context, req *link.SeeUserLinkRequest) (resp *link.SeeUserLinkResponse, err error) {
	// TODO: Your code here...
	return
}

// LinkRedirect implements the LinkServiceImpl interface.
func (s *LinkServiceImpl) LinkRedirect(ctx context.Context, req *link.LinkRedirectRequest) (resp *link.LinkRedirectResponse, err error) {
	// TODO: Your code here...
	return
}
