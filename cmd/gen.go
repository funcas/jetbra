/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/funcas/jetbra/tool"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"strings"
	"time"
)

const (
	PRI_KEY = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCElG++kNAKB/7FAYEpFR4QZ/LmqIYTnKzyBCJvacz8MUQ+cI23q/w1nvV2XLFf95QMV23yeQCCr615of26GZP0K2pPcm0EmloB9/MErBWTy5eFzffwS6RnHHJTfuGAd2d+KiM/rdoQfYvosc4kwNBdCLL4OeXeyKxeaMH/jobWMa63Do/Ljs3ZTSozVSf4sEDX6vCe8uMpgSMnAYFGTqkXSDB36hGQW7gkYxyqc2ZCFabsOjqSUWscXfGw7vJs2ZHl481hytJl87curNO+CujBDlnpXcbRKmSxv96zBbLsFmd5+EC5WRi8mRfX6lnDT4vMS+ioPwv9loeOsn+IHz6NAgMBAAECggEADv6KtM7D+J523ekRpofQ4qIfAp9EqqsplCWLj3YzzMH1qeoo3QECgT39HOnW1be2cr7gnG+68wzTXdJIU+fYp+DRITMbyU17qIVWYszjk9eOx39H71FewpqhA3hyIIlglxBE9ErBQuQwEp0pZ1a3Og5lz1Uc/llLxVGrvy9raIYERIB483nxgyYpWqiZ2kxgeLTVQJOWLOPUBmYiSW7rAhzFLcKHRsV1pQ3RjR4qZoluXSpuBpvVz/02CkwI7L/sMPd/GXURnxZ64vSsuecTMbrfiuYGsEMYj5cvvb8dWn4Q9zqkVj2V209kc1pwqEz4xnSMCNAB85jluEw5i5VlgQKBgQC4SO42Uc2B28+pWvw3jwnjskfvjouOd7us/NpNJDU/oVUMSOudOzKYMFOoLrMGJRYwN2D0ZdjFoBnPWNiHCffSfoDfisfWO8oFvlLpwXUfS0eikxK3mXJkx1dB69nZB/qIvbHfbYNUcQ+DqQKeqIUDPUSlr/3EwYG894t8asZBVQKBgQC4LHo84SH4gdp2u0R5f/Qv7uHvVOYzG0CGhzHRXiegaRvr3guuc2QaM8Sfn9HMraa8zJV335fa0a3j6lJEZk5xUODJDtH41/UgU8S34sHR/3HYbtSbOiTWPQieQ/bUuU0SXEVtj72WHt7UDclWKEYk+B4/VK3FqNkzi+UMhAhoWQKBgQCSfWbexygVkiYA+dKEV7AE0DNeunE1TvhhCp59s8qgSh7RJSYmVmhhkgjuKDEnMVPyABOSM2OdtSbPLFDzB3l4WhcXm7o/EkgFSMWgC6iOt8i58y9twwysD8nNHS6O10jQafp7IRLEjRtIDZGHjW4upQVsr/mJ7kizbhPzUu1GAQKBgQCqGlpkNENd9QFDvD7IQUfVpbsDg9SBy1/TYLbAe0f6PlpNhW4gl+8SzFGEZOYvNwvTQNmkN91Xe2YyJVAKYLt+qYaT7J4+SafbTcNZW5RMUrRhhCvVmuQ/A2wcUQV/AA6RMKvnP2WJa8W+8WTNsLDooEm/kIRXFVZW55SL9L6XMQKBgAnZSZuWnedBR1YuN710NyVPw4cGt2HnTk7hWHs6NQ/ZskuNZe/GtBrjNEnLUQr2cbaVOBtLhWifBOGNsF+v1JCVWokga/DnGGV6vNlOIqNkiuO/gcnPFyAl8R1OBRvbgK4mnsUvrb4BwKSUKM4VYIifSbgn/jP3uPxP5Ly68gYY"
	CERT    = "MIIEpTCCAo2gAwIBAgIBCTANBgkqhkiG9w0BAQsFADAYMRYwFAYDVQQDEw1KZXRQcm9maWxlIENBMB4XDTIyMDExODA3MjYzNVoXDTM3MDExODA3MjYzNVowaDELMAkGA1UEBhMCQ1oxDjAMBgNVBAgTBU51c2xlMQ8wDQYDVQQHEwZQcmFndWUxGTAXBgNVBAoTEEpldEJyYWlucyBzLnIuby4xHTAbBgNVBAMTFHByb2QzeS1mcm9tLTIwMTgxMTAxMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhJRvvpDQCgf+xQGBKRUeEGfy5qiGE5ys8gQib2nM/DFEPnCNt6v8NZ71dlyxX/eUDFdt8nkAgq+teaH9uhmT9CtqT3JtBJpaAffzBKwVk8uXhc338EukZxxyU37hgHdnfiojP63aEH2L6LHOJMDQXQiy+Dnl3sisXmjB/46G1jGutw6Py47N2U0qM1Un+LBA1+rwnvLjKYEjJwGBRk6pF0gwd+oRkFu4JGMcqnNmQhWm7Do6klFrHF3xsO7ybNmR5ePNYcrSZfO3LqzTvgrowQ5Z6V3G0Spksb/eswWy7BZnefhAuVkYvJkX1+pZw0+LzEvoqD8L/ZaHjrJ/iB8+jQIDAQABo4GpMIGmMEgGA1UdIwRBMD+AFKOetkhnQhI2Qb1t4Lm0oFKLl/GzoRykGjAYMRYwFAYDVQQDDA1KZXRQcm9maWxlIENBggkA0myxg7KDeeEwDAYDVR0TAQH/BAIwADAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDgYDVR0PAQH/BAQDAgWgMB0GA1UdDgQWBBQrcw2BwTop7cR2pYKEaIDBqtKr4zANBgkqhkiG9w0BAQsFAAOCAgEAY9EIdZPCHvHaslQwq1aLRuWjSBQNiBCikHy0bScXahxT1XIcznUltMH6OKOxZIej+r5UrLxRFeqP7rYyldecnOqKbT0QB4Buj3NYM0QakSGWCJCHuCrpvpio7fkaGM9v28ysTPDjpG5Z5QHizRyIyWbXI1yOiFq+uiiQsbZMXJoMKZohn9McSHCPy9eqj3M0Y3sj4CEU5ppjJ2/7Wt69HavN7inQ6arL+SgqqmzOuxDycbKUftYqPWmJMXqH8XeuijFea6G90WQFqz+21sIRfV3AqaQF5BxlB5Xk+PZfUo0KftwFqKXOK4UIa6Oshl6N7vN5qHKbDm0W4AOZsHSnFnKlG42gszgmpqH6w3bOo22XLU2Qk7ifarcqzshoGQYXOUHYSFHN+85fo4QPktL50p+IZRg6CNkgoBpVt5EJoUnBRqdYNxMvP/E0uudNcyiNqQCASdYDqUbpfW0wza2iRNYLp+T+EW93VZPKbXHvsbi4YT+SFwNEQyD+Uq/bvOS85smwBUA2kJ/WzWctMLSdQmwSzR0OZaA0U+GFM8hRKomREp8PM1U10ULc1NfCtCV+z6IYJ7y2M7LNp6PUOSj4ahXdQLfCmIbKTwZ/ttiIPvODeoPbIeA850UoW3T5rXRV2LIvgeHDKtBWgHU15bndFlhj3U/VhzfnRJuSkp1+krI="
)

var LIC_TEMPLATE = map[string]string{
	"ALL":          "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"II\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"DS\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"AC\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"DPN\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"RSC\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PS\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"RSF\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"GO\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"DM\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"CL\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"RS0\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"RC\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"RD\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"PC\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"RSV\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"RSU\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"RM\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"WS\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"DB\",\"paidUpTo\":\"$exp\",\"extended\":false},{\"code\":\"DC\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PDB\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PWS\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PGO\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PPS\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PPC\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PRB\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"PSW\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"DP\",\"paidUpTo\":\"$exp\",\"extended\":true},{\"code\":\"RS\",\"paidUpTo\":\"$exp\",\"extended\":true}],\"metadata\":\"0120200728EPAJA008006\",\"hash\":\"15021354/0:-1251114717\",\"gracePeriodDays\":0,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"IDEA":         "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"II\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PDB\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:1649058719\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"GOLAND":       "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"GO\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:1805249793\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"WEBSTORM":     "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"WS\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-1920204289\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"CLION":        "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"CL\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-1485202536\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"PYCHARM":      "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PC\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:738368644\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"RIDDER":       "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"RD\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PDB\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-1635216578\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"DATAGRIP":     "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"DB\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PDB\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-2014632235\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"DATASPELL":    "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"DS\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PSI\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true},{\"code\":\"PDB\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-51026298\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"RESTFUL":      "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PFASTREQUEST\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PMYBATISHELPER\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PREDIS\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:380777206\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"DOTCOVER":     "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"DC\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:1133542834\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"DOTTRACE":     "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"DPN\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"DP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":true}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-2093202651\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"DOTMEMORY":    "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"DM\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:-1534118412\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"CODE_WITH_ME": "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PCWMP\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:380777206\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"REGEX_TOOL":   "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licName\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PREGEXTOOL\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:129336052\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
	"MYBATIS":      "{\"licenseId\":\"$licId\",\"licenseeName\":\"$licNames\",\"assigneeName\":\"\",\"assigneeEmail\":\"\",\"licenseRestriction\":\"\",\"checkConcurrentUse\":false,\"products\":[{\"code\":\"PFASTREQUEST\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PMYBATISHELPER\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false},{\"code\":\"PREDIS\",\"fallbackDate\":\"$expDay\",\"paidUpTo\":\"$expDay\",\"extended\":false}],\"metadata\":\"0120230914PSAX000005\",\"hash\":\"TRIAL:380777206\",\"gracePeriodDays\":7,\"autoProlongated\":false,\"isAutoProlongated\":false}",
}

var (
	product string
	expire  string
	licName string
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate a license",
	Long:  `I am no one.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("gen called %s %s %s\n", product, expire, licName)
		generate()
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVarP(&product, "product", "p", "", "choose products")
	genCmd.Flags().StringVarP(&expire, "expire", "e", "", "expire date")
	genCmd.Flags().StringVarP(&licName, "name", "n", "", "license assign for")

	genCmd.MarkFlagRequired("product")
	genCmd.MarkFlagRequired("name")

}

func generate() {
	if len(expire) == 0 {
		year := time.Now().Year() + 10
		expire = fmt.Sprintf("%d-12-31", year)
	}
	u4 := uuid.New()
	lic := LIC_TEMPLATE[product]
	licId := strings.ToUpper(strings.ReplaceAll(u4.String(), "-", ""))[22:]
	lic = strings.Replace(lic, "$licId", licId, 1)
	lic = strings.Replace(lic, "$licName", licName, 1)
	lic = strings.ReplaceAll(lic, "$expDay", expire)
	licBase64 := base64.StdEncoding.EncodeToString([]byte(lic))
	signBase64, _ := tool.Sign(lic, PRI_KEY)
	fmt.Printf("%s-%s-%s-%s", licId, licBase64, signBase64, CERT)
}
