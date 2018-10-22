package template

const (
	// cytoscape.js 模板
	CTemp = "cs"
	// echarts.js 模板
	ETemp = "echarts"
)

var cytoscapeTemplate = []string{"<!DOCTYPE html><html><head><style type=\"text/css\">body{font:14px helvetica neue,helvetica,arial,sans-serif}#cy{height:100%;width:100%;position:absolute;left:0;top:0}</style><meta charset=utf-8/><meta name=\"viewport\" content=\"user-scalable=no, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, minimal-ui\"><title>Animated BFS</title><script src=\"https://cdnjs.cloudflare.com/ajax/libs/cytoscape/3.2.17/cytoscape.js\"></script></head><body><div id=\"cy\"></div><script>packageDate =",
	";var cy = cytoscape({container:document.getElementById('cy'),boxSelectionEnabled:false,autounselectify:true,style:cytoscape.stylesheet() .selector('node') .css({'shape':'roundrectangle','content':'data(name)','text-valign':'center','color':'white','text-outline-width':2,'text-outline-color':'#FF8C00','background-color':'#FF8C00','width':90,'height':40,'font-size':20,}) .selector('edge') .css({'curve-style':'bezier','target-arrow-shape':'triangle','line-color':'#ddd','target-arrow-color':'#ddd','arrow-scale':1.2,'width':6,}),elements:{nodes:packageDate.nodes,edges:packageDate.edges,},layout:{name:'cose',directed:true,gravity:5,padding:30,fit:false,spacingFactor:1.75,componentSpacing:50,}});</script></body></html>",
}

var echartsTemplate = []string{
	"<!DOCTYPE html><html><head><meta charset=\"UTF-8\"><title>godepView</title></head><body><div id=\"main\" style=\"width: 100%; height: 800px\"></div><script src=\"http://echarts.baidu.com/build/dist/echarts.js\"></script><script type=\"text/javascript\">packageDate =",
	";require.config({paths:{echarts:'http://echarts.baidu.com/build/dist'}});require([\"template/echarts\",\"echarts/chart/force\"],function(ec){var myChart = ec.init(document.getElementById('main'),'macarons');var option ={tooltip:{show:false},series:[{type:'force',name:\"Force tree\",itemStyle:{normal:{label:{show:true},nodeStyle:{brushType:'both',borderColor:'rgba(255,215,0,0.4)',borderWidth:1,},}},gravity:.5,linkSymbol:'arrow',useWorker:false,minRadius:15,maxRadius:25,scaling:1.1,roam:'move',categories:packageDate.categories,nodes:packageDate.nodes,links:packageDate.links,}]};myChart.setOption(option)});</script></body></html>",
}

// GetTemplate 获取页面模板
func GetTemplate(mode string) []string {
	switch mode {
	case "echarts":
		return echartsTemplate
	case "cs":
		return cytoscapeTemplate
	default:
		return cytoscapeTemplate
	}
}
