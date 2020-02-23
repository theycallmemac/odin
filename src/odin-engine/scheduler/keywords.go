package scheduler

// this function is used to return a string of keywords used
// parameters: nil
// returns: string (all valid keywords used by odin)
func getValidKeywords() string {
    return "every everyday minute hour and at 1st 2nd 3rd 4th 5th 6th 7th 8th 9th 10th 11th 12th 13th 14th 15th 16th 17th 18th 19th 20th 21st 22nd 23rd 24th 25th 26th 27th 28th 29th 30th 31st Monday Tuesday Wednesday Thursday Friday Saturday Sunday January February March April May June July August September October November December 00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59"
}

// this function is used return a map of days to numeric values
// parameters:  nil
// returns: map[string]string (mapping each valid day of the week to a number value)
func getDowMap() map[string]string {
    return map[string]string {"every Monday":"1","every Tuesday":"2","every Wednesday":"3","every Thursday":"4","every Friday":"5","every Saturday":"6","every Sunday":"7","everyday": "*",}
}

// this function is used to return a map of days of the month to numeric values
// parameters:  nil
// returns: map[string]string (mapping each valid day of the month to a number value)
func getDomMap() map[string]string {
    return map[string]string {"every 1st":"1","every 2nd":"2","every 3rd":"3","every 4th":"4","every 5th":"5","every 6th":"6","every 7th":"7","every 8th":"8","every 9th":"9","every 10th":"10","every 11th":"11","every 12th":"12","every 13th":"13","every 14th":"14","every 15th":"15","every 16th":"16","every 17th":"17","every 18th":"18","every 19th":"19","every 20th":"20","every 21st":"21","every 22nd":"2","every 23rd":"23","every 24th":"24","every 25th":"25","every 26th":"26","every 27th":"27","every 28th":"28","every 29th":"29","every 30th":"30","every 31st":"31",}
}

// this function is used to return a map of months of the year to numeric values
// parameters:  nil
// returns: map[string]string (mapping each valid month of the year to a number value)
func getMonMap() map[string]string {
    return map[string]string {"every January":"1","every February":"2","every March":"3","every April":"4","every May":"5","every June":"6","every July":"7th","every August":"8","every September": "9","every October":"10","every November":"11","every December":"12",}
}
