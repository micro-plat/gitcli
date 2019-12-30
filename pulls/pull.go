package pulls

import (
	"fmt"
	"github.com/micro-plat/cli/cmds"
	"github.com/micro-plat/cli/logs"
	"github.com/micro-plat/gitcli/gitlabs"
	"github.com/micro-plat/lib4go/types"
	"github.com/urfave/cli"
)

func init() {
	cmds.Register(
		cli.Command{
			Name:  "pull",
			Usage: "拉取最新",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "branch,b",
					Usage: "分支",
				},
			},
			Action: pull,
		})
}

//pull 根据传入的路径(分组/仓库)拉取所有项目
func pull(c *cli.Context) (err error) {
	branch := types.GetString(c.String("branch"), "master")
	reps, err := gitlabs.GetRepositories(c.Args().Get(0))
	if err != nil {
		return err
	}
	if len(reps) == 0 {
		return fmt.Errorf("没有需要拉取的项目")
	}
	for _, rep := range reps {
		if !rep.Exists() {
			logs.Log.Infof("get clone %s %s", rep.FullPath, rep.GetLocalPath())
			if err := rep.Clone(); err != nil {
				logs.Log.Error(err)
			}
		}
		if err := rep.Pull(branch); err != nil {
			logs.Log.Error(err)
		}
	}
	return nil

}
