package getopt

import (
	"fmt"
	"os"
	"strings"
)

type Args struct {
	program string
	params  []string
}

var expectedArgs = newSet()

func (a *Args) GetProgram() string {
	return a.program
}

func (a *Args) GetParams() []string {
	return a.params
}

func AddFlag(shortName rune, longName string, p *bool) {
	f := &flag{
		option: option{
			shortName: shortName,
			longName:  longName,
		},
		p: p,
	}
	expectedArgs.expectedShort[shortName] = f
	expectedArgs.expectedLong[longName] = f
}

func AddOption(shortName rune, longName string, p interface{}, isOptional bool, valueParser ValueParser) {
	o := &valuedOption{
		option: option{
			shortName: shortName,
			longName:  longName,
		},
		value:       p,
		valueParser: valueParser,
		isOptional:  isOptional,
	}
	expectedArgs.expectedShort[shortName] = o
	expectedArgs.expectedLong[longName] = o
}

func Parse() (a *Args, err error) {
	return parse(os.Args)
}

func parse(args []string) (a *Args, err error) {
	a = &Args{
		program: args[0],
	}

	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg[0] == '-' {
			if arg[1] == '-' {
				kv := strings.IndexRune(arg, '=')
				var value string
				if kv > 0 {
					value = arg[kv+1:]
					arg = arg[:kv]
				}
				opt := expectedArgs.expectedLong[arg[2:]]
				if opt != nil {
					flag, ok := opt.(*flag)
					if ok {
						flag.set()
					} else {
						valueOpt, ok := opt.(*valuedOption)
						if !ok {
							a.params = append(a.params, arg)
							continue
						}
						if value == "" && !valueOpt.isOptional && kv < 0 {
							if len(args) < i+2 {
								err = fmt.Errorf("missing argument value")
								return
							}
							i++
							value = args[i]
						}
						err = valueOpt.setValue(value)
						if err != nil {
							return
						}
					}
				}
			} else {
				for _, r := range arg[1:] {
					opt := expectedArgs.expectedShort[r]
					if opt != nil {
						flag, ok := opt.(*flag)
						if ok {
							flag.set()
						} else {
							valueOpt, ok := opt.(*valuedOption)
							if !ok {
								a.params = append(a.params, arg)
								continue
							}
							if !valueOpt.isOptional && len(args) < i+2 {
								err = fmt.Errorf("missing argument value")
								return
							}
							i++
							err = valueOpt.setValue(args[i])
							if err != nil {
								return
							}
							break
						}
					}
				}
			}
		} else {
			a.params = append(a.params, arg)
		}
	}
	return
}
