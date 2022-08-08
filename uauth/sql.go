/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>
 */

package uauth

var sqlUserActive = "UPDATE user_account set active=$1 WHERE user_account_id=$2;"

var sqlDeleteUserAccount = "DELETE FROM user_account WHERE user_account_id=$1;"

var sqlChangeRole = "UPDATE user_account set role=$1 WHERE user_account_id=$2;"

var sqlCreateSession = "INSERT INTO user_session" +
	"(user_account_id, auth_token, session_expiry)" +
	"VALUES ($1,$2,$3);"

var sqlFetchAuthUser = "SELECT user_session.session_expiry, " +
	"user_account.full_name,user_account.email,user_account.phone,user_account.user_account_id,user_account.active,user_account.role " +
	"from user_session, user_account WHERE auth_token=$1 " +
	"AND user_session.user_account_id=user_account.user_account_id;"

var sqlFetchUsers = "SELECT * FROM user_account LIMIT $1 OFFSET $2;"
var sqlTotalUsers = "SELECT count(*) FROM user_account;"
var sqlDeleteUser = "DELETE from user_account WHERE user_account_id=$1;"
var sqlChangePassword = "UPDATE user_account set password=$1 WHERE user_account_id=$2;"

var sqlSearchUsers = "SELECT * FROM user_account WHERE full_name ILIKE $1 OR email ILIKE $1 OR phone ILIKE $1 " +
	"LIMIT $2 OFFSET $3;"

var sqlSearchUsersTotal = "SELECT count(*) FROM user_account WHERE full_name LIKE $1 OR email LIKE $1 OR phone LIKE $1;"
var sqlDeletePasswordResetKey = "delete from user_keys where user_account_id=$1 and key_name=$2;"
