package models

import "context"

type Client struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (c *Client) GetAllClients() ([]*Client, error) {
	// Ctx so it closes after 3 scs
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()

	query := `select * from client`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var clients []*Client
	for rows.Next() { // save each row
		var client Client
		err := rows.Scan( // save each parameter in client
			&client.ID,
			&client.FirstName,
			&client.LastName,
		)
		if err != nil {
			return nil, err
		}
		clients = append(clients, &client) // im speechless by this way of appending
	}
	return clients, nil
}
