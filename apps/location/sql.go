/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, August 2019
 */

package location

var sqlSelectCountry = "select * from country;"
var sqlSelectState = "select * from state where country_id=$1;"
var sqlSelectDistrict = "select * from district where state_id=$1 order by district_name;"
var sqlSelectCity = "select city_id,city_name,district_id from city where district_id=$1 order by city_name;"
var sqlSelectLocalBodyType = "select * from local_body_type where country_id=$1;"
var sqlSelectLocalBody = "select * from local_body where district_id=$1 and local_body_type_id=$2;"
