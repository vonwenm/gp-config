// gp-config - A configuration file parser for golang.
//
// 1. Overview
//
// gp-config is a configuration file parser for go that loosely follows the
// TOML syntax.
//
// Main differences are:
//
// * section and option names are case insensitive;
// * multi-dimensional arrays are not supported;
// * tables are not supported; and
// * array of tables are not supported.
//
// Why? I just don't need multi-dimensional arrays or tables in my configuration
// files.
//
// Parser implements following grammar (EBNF style):
//
// 	  config = section | options
// 	  section = '[' IDENTIFIER ']' EOL options
// 	  options = option {option}
// 	  option =  IDENTIFIER '=' (value | array) EOL
//    value = BOOL | INT | FLOAT | DATE | STRING
//    array = '[' {EOF} value {EOF} {, {EOF} value {EOF} } ']'
//
// 2. Usage
//
// 2.1. Loading Configuration Files
//
// Configuration can either be stored in a string or a file.
//
//    import (
//        "flag"
//        "fmt"
//        "github.com/cbonello/gp-config"
//        "os"
//    )
//
//    const deflt = `version = [1, 0, 10]
//                   [database]
//                       dbname = "mydb"
//                       user = "foo"
//                       password = "bar"`
//
//    var dev bool = false
//
//    func main() {
//        flag.BoolVar(&dev, "dev", false, "Runs application in debug mode, default is production.")
//        flag.Parse()
//        // Set default options (Production mode).
//        cfg := config.NewConfiguration()
//        if err := cfg.LoadString(deflt); err != nil {
//            fmt.Printf("error: default config: %d:%d: %s\n",
//            err.Line, err.Column, err)
//            os.Exit(1)
//        }
//        if dev {
//            // Override default options with debug mode settings.
//            if err := cfg.LoadFile("debug.cfg"); err != nil {
//                fmt.Printf("error: %s: %d:%d: %s\n",
//                err.Filename, err.Line, err.Column, err)
//                os.Exit(1)
//            }
//        }
//        ...
//    }
//
// Contents of `debug.cfg` may for instance be:
//
//    [database]
//        dbname = "mydb_test"
//
// 2.2. Reading Configuration Files
//
// 2.2.1. Basic API
//
//    if dbname, err := config.GetString("database.dbname"); err != nil {
//        fmt.Printf("error: configuration: missing database name\n")
//        os.Exit(1)
//    }
//    user := config.GetStringDefault("database.dbname", "user")
//    password := config.GetStringDefault("database.dbname", "foobar")
//
//
// 2.2.2. Reflection API
//
// Following structure may be declared to record options of `[database]`
// section defined above.
//
//    type database struct {
//        Name     string `option:"dbname"`
//        Username string `option:"user"`
//        Password string
//    }
//
// A structure field annotation may be used if there is no direct match
// between a field name and an option name.
//
// And finally, to decode:
//
//    var db database
//        if err := cfg.Decode("server", &db); err != nil {
//            fmt.Println("error:", err)
//            os.Exit(1)
//        }
//
// 3. Examples
//
// Demo applications are provided in the `examples/` directory. To launch
// them:
//
//    go run github.com/cbonello/gp-config/examples/demo/demo.go
//    go run github.com/cbonello/gp-config/examples/demo-decode/demo.go
//

package config
