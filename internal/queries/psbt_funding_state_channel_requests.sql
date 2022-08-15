-- name: SetPsbtFundingStateChannelRequest :exec
INSERT INTO psbt_funding_state_channel_requests (psbt_funding_state_id, channel_request_id)
  VALUES ($1, $2);

-- name: ListPsbtFundingStateChannelRequests :many
SELECT cr.* FROM channel_requests cr
  INNER JOIN psbt_funding_state_channel_requests pfscr ON pfscr.channel_request_id = cr.id
  WHERE pfscr.psbt_funding_state_id = $1
  ORDER BY cr.id;
