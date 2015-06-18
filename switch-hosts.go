package main

import(
  "os"
  "os/exec"
  "fmt"
  "strings"
  "io/ioutil"
  homedir "github.com/mitchellh/go-homedir"
)

func printUsage() {
  fmt.Println("Usage of switch-hosts:")
  fmt.Println("switch-hosts <HOSTS_CONFIG>")
  fmt.Println("<HOSTS_CONFIG> must be a file present in ~/.hosts/")
  fmt.Println("Use the 'default' config to restore the original hosts file")
}

func hasHostsDir() bool {
  hosts_dir, _ := homedir.Expand("~/.hosts/")
  stat, err := os.Stat(hosts_dir)

  if err != nil {
    return false
  }

  return stat.IsDir()
}

func createHostsDir() {
  dir, _ := homedir.Expand("~/.hosts/")
  execCmd("mkdir", "-p", dir)
}

func hostConfigPath(config string) string {
  path := fmt.Sprintf("~/.hosts/%s", config)
  expanded, _ := homedir.Expand(path)
  return expanded
}

func fileExist(filename string) bool {
  _, err := os.Stat(filename)
  return err == nil
}

func hostConfigExist(config string) bool {
  path := hostConfigPath(config)
  return fileExist(path)
}

func applyConfig(config string) {
  configPath := hostConfigPath(config)

  if config == "default" {
    execCmd("sudo", "cp", configPath, "/etc/hosts")
  } else {
    createAndApplyConfig(config)
  }
}

func createAndApplyConfig(config string) {
  defaultFileContents, _ := ioutil.ReadFile(hostConfigPath("default"))
  configFileContents, _ := ioutil.ReadFile(hostConfigPath(config))

  fileContents := append(defaultFileContents, byte('\n'))
  fileContents = append(fileContents, configFileContents...)
  tmpFile, _ := ioutil.TempFile("", "hosts-")
  defer tmpFile.Close()

  ioutil.WriteFile(tmpFile.Name(), fileContents, os.ModeTemporary)
  execCmd("sudo", "cp", tmpFile.Name(), "/etc/hosts")
  os.Remove(tmpFile.Name())
}

func hasHostsBackup() bool {
    return hostConfigExist("default")
}

func createHostsBackup() {
  if fileExist("/etc/hosts.orig") {
    execCmd("cp", "/etc/hosts.orig", hostConfigPath("default"))
  } else {
    execCmd("cp", "/etc/hosts", hostConfigPath("default"))
    execCmd("sudo", "cp", "/etc/hosts", "/etc/hosts.orig")
  }
}

func execCmd(command string, args ...string) string {
  cmd := exec.Command(command, args...)
  stdout, _ := cmd.StdoutPipe()

  cmd.Start()

  output := make([]byte, 1024)
  bytesRead, _ := stdout.Read(output)

  cmd.Wait()

  output = output[0:bytesRead]

  return strings.TrimSpace(string(output))
}


func main() {
  if !hasHostsDir() {
    createHostsDir()
  }

  if !hasHostsBackup() {
    fmt.Println("No hosts file backup found... creating")
    createHostsBackup()
  }

  if len(os.Args) == 1 || strings.TrimSpace( os.Args[1] ) == ""  {
    printUsage();
    return
  }

  config := strings.TrimSpace( os.Args[1] )

  if hostConfigExist(config) {
    applyConfig(config)
    fmt.Printf("Hosts configuration switched to %s.\n", config)
  } else {
    fmt.Printf("%s hosts configuration was not found. Did you define it at ~/.hosts/%s ?\n\n", config, config)
    printUsage()
    return
  }
}
