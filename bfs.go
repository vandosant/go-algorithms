package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func GetPath(g map[string][]string, s string) map[string]map[string]int {
    bfsData := map[string]map[string]int{}
    
    for k, _ := range g {
        m := map[string]int{
            "distance":-1,
            "predecessor":-1,
        }
        bfsData[k] = m
    }
    
    bfsData[s]["distance"] = 0
    
    q := []string{}
    q = append(q, s)
    
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        for i := 0; i < len(g[u]); i++ {
            v := g[u][i]
            if bfsData[v]["distance"] == -1 || bfsData[v]["distance"] > bfsData[u]["distance"] + 6 {
                bfsData[v]["distance"] = bfsData[u]["distance"] + 6
                p, err := strconv.Atoi(u)
                if err != nil {
                    panic(err)
                }
                bfsData[v]["predecessor"] = p
                q = append(q, v)
            }
        }
    }
    
    return bfsData
}

func main() {  
    input := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }
    e := 0
    for i := 1; i < len(input); i+=e+2 {
        test_graph := map[string][]string{}
        this_info := strings.Fields(input[i])
        n, err := strconv.Atoi(this_info[0])
        if err != nil {
            panic(err)
        }
        this_e, err := strconv.Atoi(this_info[1])
        if err != nil {
            panic(err)
        }
        e = this_e

        for k := 1; k <= n; k++ {
            v := []string{}
            test_graph[strconv.Itoa(k)] = v
        }

        this_graph := input[i+1:i+e+1]
        start_node := strings.TrimSpace(input[i+e+1])
        for _, j := range this_graph {
            nodes := strings.Fields(j)
            test_graph[nodes[0]] = append(test_graph[nodes[0]], nodes[1])
        }

        d := GetPath(test_graph, start_node)
        for i := 0; i < len(d); i++ {
            if (strconv.Itoa(i + 1) != start_node) {
                fmt.Printf("%d ", d[strconv.Itoa(i + 1)]["distance"])
            }
        }
        fmt.Printf("\n")
    }
}
