/*
 * Copyright (C) Centipair Technologies Private Limited - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 * Written By Devasia Joseph <devasiajoseph@centipair.com>, August 2018
 */

package cpmath

func LimitOffset(page int, count int) (int, int) {
	limit := count
	if count <= 0 {
		limit = 50
	}

	offset := 0
	if page <= 1 {
		return limit, offset
	} else {
		offset = (page - 1) * count
		return limit, offset
	}

}

func Offset(page int, count int) int {
	offset := 0
	if page <= 1 {
		return offset
	} else {
		offset = (page - 1) * count
		return offset
	}

}
