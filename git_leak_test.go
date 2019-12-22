#go-test
#这是一条测试，请忽略

func main() {

    params := url.Values{}

    Url, err := url.Parse("http://gitlab.hutaojie.com")
    if err != nil {
        panic(err.Error())

    }
    params.Set("git_leak_test", "git_leak_test")
    params.Set("id", string("1"))
    Url.RawQuery = params.Encode()
    urlPath := Url.String()
    resp, err := http.Get(urlPath)
    defer resp.Body.Close()
    s, err := ioutil.ReadAll(resp.Body)
    fmt.Println(string(s))

}
