/*
Copyright (c) 2022, Vaclovas Lapinskis. All rights reserved.
*/

package adminserver

import (
	"context"

	"webimizer.dev/poem/admin"
)

func (srv *adminServer) AddPoem(_ context.Context, poem *admin.AdminPoem) (response *admin.PoemResponse, err error) {
	err = srv.cmd.addPoem(poem)
	if err != nil {
		return nil, err
	}
	response = &admin.PoemResponse{Success: true, Poem: poem}
	return response, nil
}

func (p *adminServerCmd) addPoem(poem *admin.AdminPoem) error {
	db, err := p.openDBConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO `poem_poems`(category_id,title,text) VALUES (?,?,?);", poem.CategoryId, poem.Title, poem.Text)
	if err != nil {
		return err
	}
	return nil
}
