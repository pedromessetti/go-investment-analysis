package main

import (
	"log"
	"time"
	"investment-analysys/internal/scraper"
)

func getDate() string {
	return time.Now().Format("2006-01-02")
}

func main() {
	FundamentusScraper := scraper.Scraper{
		URL: "https://www.fundamentus.com.br/resultado.php",
		SourceName: "Fundamentus",
		Date: getDate(),
	}

	StatusInvestScraper := scraper.Scraper{
		URL: "https://statusinvest.com.br/category/AdvancedSearchResultExport?search=%7B%22Sector%22%3A%22%22%2C%22SubSector%22%3A%22%22%2C%22Segment%22%3A%22%22%2C%22my_range%22%3A%22-20%3B100%22%2C%22forecast%22%3A%7B%22upsidedownside%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22estimatesnumber%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22revisedup%22%3Atrue%2C%22reviseddown%22%3Atrue%2C%22consensus%22%3A%5B%5D%7D%2C%22dy%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_l%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22peg_ratio%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_vp%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_ativo%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22margembruta%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22margemebit%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22margemliquida%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_ebit%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22ev_ebit%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22dividaliquidaebit%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22dividaliquidapatrimonioliquido%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_sr%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_capitalgiro%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22p_ativocirculante%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22roe%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22roic%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22roa%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22liquidezcorrente%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22pl_ativo%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22passivo_ativo%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22giroativos%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22receitas_cagr5%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22lucros_cagr5%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22liquidezmediadiaria%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22vpa%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22lpa%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%2C%22valormercado%22%3A%7B%22Item1%22%3Anull%2C%22Item2%22%3Anull%7D%7D&CategoryType=1",
		SourceName: "StatusInvest",
		Date: getDate(),
	}

	// InvestSiteScraper := scraper.Scraper{
	// 	URL: "https://www.investsite.com.br/selecao_acoes.php?atual%2522%255D&todos=todos&ROTanC_min=&ROTanC_max=&chk_lst%5B%5D=itm7&ROInvC_min=&ROInvC_max=&chk_lst%5B%5D=itm8&ROE_min=&ROE_max=&chk_lst%5B%5D=itm9&ROA_min=&ROA_max=&chk_lst%5B%5D=itm10&margem_liquida_min=&margem_liquida_max=&chk_lst%5B%5D=itm11&margem_bruta_min=&margem_bruta_max=&chk_lst%5B%5D=itm12&margem_EBIT_min=&margem_EBIT_max=&chk_lst%5B%5D=itm13&giro_ativo_min=&giro_ativo_max=&chk_lst%5B%5D=itm14&fin_leverage_min=&fin_leverage_max=&chk_lst%5B%5D=itm15&debt_equity_min=&debt_equity_max=&chk_lst%5B%5D=itm16&p_e_min=&p_e_max=&chk_lst%5B%5D=itm17&p_bv_min=&p_bv_max=&chk_lst%5B%5D=itm18&p_receita_liquida_min=&p_receita_liquida_max=&chk_lst%5B%5D=itm19&p_FCO_min=&p_FCO_max=&chk_lst%5B%5D=itm20&p_FCF1_min=&p_FCF1_max=&chk_lst%5B%5D=itm21&p_EBIT_min=&p_EBIT_max=&chk_lst%5B%5D=itm22&p_ncav_min=&p_ncav_max=&chk_lst%5B%5D=itm23&p_ativo_total_min=&p_ativo_total_max=&chk_lst%5B%5D=itm24&p_capital_giro_min=&p_capital_giro_max=&chk_lst%5B%5D=itm25&EV_EBIT_min=&EV_EBIT_max=&chk_lst%5B%5D=itm26&EV_EBITDA_min=&EV_EBITDA_max=&chk_lst%5B%5D=itm27&EV_receita_liquida_min=&EV_receita_liquida_max=&chk_lst%5B%5D=itm28&EV_FCO_min=&EV_FCO_max=&chk_lst%5B%5D=itm29&EV_FCF1_min=&EV_FCF1_max=&chk_lst%5B%5D=itm30&EV_ativo_total_min=&EV_ativo_total_max=&chk_lst%5B%5D=itm31&div_yield_min=&div_yield_max=&chk_lst%5B%5D=itm32&vol_financ_min=&vol_financ_max=&chk_lst%5B%5D=itm33&market_cap_min=&market_cap_max=&chk_lst%5B%5D=itm34&setor=",
	// 	SiteName: "InvestSite",
	// 	Date: helpers.GetDate(),
	// }

	_, err := scraper.GetFromURL(StatusInvestScraper.URL)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sucessfully get data from %s", StatusInvestScraper.SourceName)
	
	_, err = scraper.ScrapeTableFromURL(FundamentusScraper)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sucessfully scraped data from %s", FundamentusScraper.SourceName)
}
