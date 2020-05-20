package main

import (
    "fmt"
    "strconv"
    "math"
)

// Converts the scanned tag number into the format
// the board itself wants
func convertTagNum(tagSN int) (string, error) {
    // Convert the tag to binary and make sure
    // there are 24 bits
    bins := strconv.FormatInt(int64(tagSN), 2)
    bins = fmt.Sprintf("%024s", bins)   

    // Now split into the facility code and user code
    frontb := bins[0:8]
    backb := bins[len(bins)-16:]

    //fmt.Println("\t - frontb", frontb)
    //fmt.Println("\t - backb", backb)

    f, err := strconv.ParseInt(frontb, 2, 32)
    if err != nil {
            return "", err
    }
    //fmt.Println("\tf:", f)

    b, err := strconv.ParseInt(backb, 2, 32)    
    if err != nil {
            return "", err
    }
    //fmt.Println("\tb:", b)
    
    return fmt.Sprintf("%d%d", f, b), nil
}

func convertBoardNum(boardTag string) (string, error) { 
    // First we want to split our string into two parts, the 
    // last five digits are our user code, all the digits 
    // in before that are the facility code. Note that there
    // could be three or less digits for the facility code,
    // which is why we have to be careful about splitting
    // into the two parts
    facilityCode := boardTag[0:len(boardTag)-5]
    userCode := boardTag[len(boardTag)-5:]
    
    //fmt.Printf("FC: %s, UC: %s\n", facilityCode, userCode)
    
    fcNum, err := strconv.Atoi(facilityCode)
    if err != nil {
        fmt.Println("Hmm, got", err)
        return "", err
    }

    fcBins := strconv.FormatInt(int64(fcNum), 2)
    
    // This is the number of bits we working with, which
    // we work down to 0
    bitCountdown := 24
    
    //////////////////////////////////////////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////
    // Facility Code
    //////////////////////////////////////////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////

    // The facility code may be less than three digits, in which case
    // our bits won't be the right number. The solution is to pad the front
    // of the slice with zeros
    fcBins = fmt.Sprintf("%08v", fcBins)
    //fmt.Println("FC Binary\t", fcBins)

    // Now let's go through the bits in the facility code table
    var fcBitTable [8]int
    fcIdx := 0  
    for _, char := range fcBins {       
        testBit := string(char)     
        if testBit == "1" {         
            fcBitTable[fcIdx] = int(math.Pow(2, float64(bitCountdown - 1)))
        } else {
            fcBitTable[fcIdx] = 0
        }
        
        fcIdx++
        bitCountdown--
    }

    fcSum := 0
    for _, num := range fcBitTable {
        //fmt.Println("-->", num)
        fcSum += num
    }


    //////////////////////////////////////////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////
    // User Code
    //////////////////////////////////////////////////////////////////////////
    //////////////////////////////////////////////////////////////////////////
    ucNum, err := strconv.Atoi(userCode)
    if err != nil {
        fmt.Println("Hmm, got", err)
        return "", err
    }
    ucBins := strconv.FormatInt(int64(ucNum), 2)
    ucBins = fmt.Sprintf("%16v", ucBins)
    //fmt.Println("User binary\t", ucBins)

    var ucBitTable [16]int
    ucIdx := 0  
    for _, char := range ucBins {       
        testBit := string(char)     
        if testBit == "1" {         
            ucBitTable[ucIdx] = int(math.Pow(2, float64(bitCountdown - 1)))
        } else {
            ucBitTable[ucIdx] = 0
        }
        
        ucIdx++
        bitCountdown--
    }

    ucSum := 0
    for _, num := range ucBitTable {
        //fmt.Println("==>", num)
        ucSum += num
    }

    return strconv.Itoa(fcSum + ucSum), nil
}

func main() {

    // Start with the user's tag that would be
    // in his or her possession. This is sometimes
    // written on the card or tag, otherwise you
    // need to use a reader to get it. Note that
    // if the number has zeros at the front, those
    // should be removed
    tagNum := 5524050

    // Convert the user's tag into the value the
    // board will expect it
    boardTag, err := convertTagNum(tagNum)
    if err != nil {
        fmt.Println("Hmm, got", err)
    }
    fmt.Printf("(To the board) %d is now %s\n", tagNum, boardTag)

    fmt.Println("=======")
    
    // Now let's try to convert it back...  
    userTag, err := convertBoardNum(boardTag)
    if err != nil {
        fmt.Println("Hmm, got", err)
    }
    fmt.Printf("(From the board) %s is now %s\n", boardTag, userTag)
}
