package main
import "io/ioutil"
import "fmt"
import "math"
import "strings"
import "strconv"

const path string = "input"

func main() {

    //Get my file
    file, _ := ioutil.ReadFile(path)
    q1Sum := 0
    q2Sum := 0
    //each line
    for  _, row := range strings.Split(string(file), "\n") {
        min, max := 999999, -99999
        var currentRow []int
        if row != "" {
        //for each int
            for _,current := range strings.Split(row," ") {
                if current != "" {
                    tmp,_ := strconv.Atoi(current)
                    if tmp < min {
                       min = tmp
                    }
                    if tmp > max {
                        max = tmp
                    }
                    for _,exist := range currentRow {
                        if exist < tmp {
                            if math.Mod(float64(tmp),float64(exist)) == 0 {
                                q2Sum += tmp / exist
                            }
                        } else {
                            if math.Mod(float64(exist),float64(tmp)) == 0 {
                                q2Sum += exist / tmp
                            }
                        }
                    }
                    currentRow = append(currentRow,tmp)
                }
            }
            q1Sum += max-min
        }
    }
    fmt.Print("The sum of the min max of each line is: ")
    fmt.Println(q1Sum)
    fmt.Print("The sum of the dividable of each line is: ")
    fmt.Println(q2Sum)
}
