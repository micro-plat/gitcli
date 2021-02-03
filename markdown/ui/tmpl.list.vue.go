package ui

const TmplList = `
{{- $string := "string" -}}
{{- $int := "int" -}}
{{- $int64 := "int64" -}}
{{- $decimal := "types.Decimal" -}}
{{- $time := "time.Time" -}}
{{- $len := 32 -}}
{{- $rows := .Rows -}}
{{- $pks := .|pks -}}
{{- $tb :=. -}}
<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
			{{- range $i,$c:=$rows|query}}
				{{- if $c.Con|TA}}
				<el-form-item>
					<el-input type="textarea" :rows="2" placeholder="请输入{{$c.Desc|shortName}}" v-model="queryData.{{$c.Name}}">
					</el-input>
				</el-form-item>
				{{- else if $c.Con|RD}}
				<el-form-item  label="{{$c.Desc|shortName}}:">
					<el-radio-group v-model="queryData.{{$c.Name}}" style="margin-left:5px">
            <el-radio v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :label="item.value">{{"{{item.name}}"}}</el-radio>
          </el-radio-group>
				</el-form-item>
				{{- else if $c.Con|SL }}
				<el-form-item>
					<el-select size="medium" v-model="queryData.{{$c.Name}}" class="input-cos" placeholder="请选择{{$c.Desc|shortName}}">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
        {{- else if or ($c.Con|qCon|DTP) (and (not ($c.Con|qCon|DP)) ($c.Con|DTP)) }}
					<el-form-item label="{{$c.Desc|shortName}}:">
						<el-date-picker class="input-cos" v-model="{{$c.Name|lowerName}}" type="datetime" value-format="yyyy-MM-dd HH:mm:ss"  placeholder="选择日期"></el-date-picker>
					</el-form-item>
				{{- else if or ($c.Con|qCon|DP) (and (not ($c.Con|qCon|DTP)) ($c.Con|DP)) }}
					<el-form-item label="{{$c.Desc|shortName}}:">
						<el-date-picker class="input-cos" v-model="{{$c.Name|lowerName}}" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
					</el-form-item>
				{{- else if $c.Con|CB }}
				<el-form-item label="{{$c.Desc|shortName}}:">
          <el-checkbox-group v-model="queryData.{{$c.Name}}">
          	<el-checkbox v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :value="item.value" :label="item.name"></el-checkbox>
          </el-checkbox-group>
        </el-form-item>
				{{- else}}
				<el-form-item>
					<el-input clearable v-model="queryData.{{$c.Name}}" placeholder="请输入{{$c.Desc|shortName}}">
					</el-input>
				</el-form-item>
				{{- end}}
			{{end}}
				{{- if gt ($rows|query|len) 0}}
				<el-form-item>
					<el-button type="primary" @click="query" size="small">查询</el-button>
				</el-form-item>
				{{end}}
				{{- if gt ($rows|create|len) 0}}
				<el-form-item>
					<el-button type="success" size="small" @click="showAdd">添加</el-button>
				</el-form-item>
				{{end}}
			</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="dataList.items" border style="width: 100%">
				{{- range $i,$c:=$rows|list}}
				<el-table-column prop="{{$c.Name}}" label="{{$c.Desc|shortName}}" align="center">
				{{- if or ($c.Con|SL) ($c.Con|CB) ($c.Con|RD)}}
					<template slot-scope="scope">
						<span {{if ($c.Con|CC)}}:class="scope.row.{{$c.Name}}|fltrTextColor"{{end}}>{{"{{scope.row."}}{{$c.Name}} | fltrEnum("{{(or (dicType $c.Con $tb) $c.Name)|lower}}")}}</span>
					</template>
				{{- else if and (eq ($c.Type|codeType) $string) (gt $c.Len $len )}}
					<template slot-scope="scope">
						<el-tooltip class="item" v-if="scope.row.{{$c.Name}} && scope.row.{{$c.Name}}.length > 20" effect="dark" placement="top">
							<div slot="content" style="width: 110px">{{"{{scope.row."}}{{$c.Name}}}}</div>
							<span>{{"{{scope.row."}}{{$c.Name}} | fltrSubstr(20) }}</span>
						</el-tooltip>
						<span v-else>{{"{{scope.row."}}{{$c.Name}}}}</span>
					</template>
				{{- else if and (or (eq ($c.Type|codeType) $int64) (eq ($c.Type|codeType) $int)) (ne $c.Name ($pks|firstStr))}}
				<template slot-scope="scope">
					<span>{{"{{scope.row."}}{{$c.Name}} | fltrNumberFormat({{or ($c.Con|decimalCon) "0"}})}}</span>
				</template>
				{{- else if eq ($c.Type|codeType) $decimal }}
				<template slot-scope="scope">
					<span>{{"{{scope.row."}}{{$c.Name}} | fltrNumberFormat({{or ($c.Con|decimalCon) "2"}})}}</span>
				</template>
				{{- else if eq ($c.Type|codeType) $time }}
				<template slot-scope="scope">
					<span>{{"{{scope.row."}}{{$c.Name}} | {{if or ($c.Con|lCon|DTP) (and (not ($c.Con|lCon|DP)) ($c.Con|DTP)) }}{{($c.Con|lCon)}}fltrDate("yyyy-MM-dd hh:mm:ss"){{else}}fltrDate{{end}} }}</span>
				</template>
				{{- else}}
				<template slot-scope="scope">
					<span>{{"{{scope.row."}}{{$c.Name}}}}</span>
				</template>
				{{end}}
				</el-table-column>
				{{- end}}
				<el-table-column  label="操作">
					<template slot-scope="scope">
						{{- if gt ($rows|update|len) 0}}
						<el-button type="text" size="small" @click="showEdit(scope.row)">编辑</el-button>
						{{- end}}
						{{- if gt ($rows|detail|len) 0}}
						<el-button type="text" size="small" @click="showDetail(scope.row)">详情</el-button>
						{{- end}}
						{{- if gt ($rows|delete|len) 0}}
						<el-button type="text" size="small" @click="del(scope.row)">删除</el-button>
						{{- end}}
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		{{if gt ($rows|create|len) 0 -}}
		<!-- Add Form -->
		<Add ref="Add" :refresh="query"></Add>
		<!--Add Form -->
		{{- end}}

		{{if gt ($rows|update|len) 0 -}}
		<!-- edit Form start-->
		<Edit ref="Edit" :refresh="query"></Edit>
		<!-- edit Form end-->
		{{- end}}

		<!-- pagination start -->
		<div class="page-pagination">
		<el-pagination
			@size-change="pageSizeChange"
			@current-change="pageIndexChange"
			:current-page="paging.pi"
			:page-size="paging.ps"
			:page-sizes="paging.sizes"
			layout="total, sizes, prev, pager, next, jumper"
			:total="dataList.count">
		</el-pagination>
		</div>
		<!-- pagination end -->

	</div>
</template>


<script>
{{- if gt ($rows|create|len) 0}}
import Add from "./{{.Name|rmhd|l2d}}.add"
{{- end}}
{{- if gt ($rows|update|len) 0}}
import Edit from "./{{.Name|rmhd|l2d}}.edit"
{{- end}}
export default {
  components: {
		{{- if gt ($rows|create|len) 0}}
		Add,
		{{- end}}
		{{- if gt ($rows|update|len) 0}}
		Edit
		{{- end}}
  },
  data () {
		return {
			paging: {ps: 10, pi: 1,total:0,sizes:[5, 10, 20, 50]},
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
      queryData:{},               //查询数据对象 
			{{- range $i,$c:=$rows|query -}}
			{{if or ($c.Con|SL) ($c.Con|CB) ($c.Con|RD) }}
			{{$c.Name|lowerName}}: this.$enum.get("{{(or (dicType $c.Con $tb) $c.Name)|lower}}"),
			{{- end}}
			{{- if or ($c.Con|qCon|DTP) (and (not ($c.Con|qCon|DP)) ($c.Con|DTP)) }}
			{{$c.Name|lowerName}}: this.$utility.dateFormat(new Date(),"yyyy-MM-dd 00:00:00"),
			{{- else if or ($c.Con|qCon|DP) (and (not ($c.Con|qCon|DTP)) ($c.Con|DP))  }}
			{{$c.Name|lowerName}}: this.$utility.dateFormat(new Date(),"yyyy-MM-dd"),{{end}}
      {{- end}}
			dataList: {count: 0,items: []}, //表单数据对象
		}
  },
  created(){
  },
  mounted(){
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
    query(){
      this.queryData.pi = this.paging.pi
			this.queryData.ps = this.paging.ps
			{{- range $i,$c:=$rows|query -}}
			{{- if or ($c.Con|qCon|DTP) (and (not ($c.Con|qCon|DP)) ($c.Con|DTP)) }}
			this.queryData.{{$c.Name}} = this.$utility.dateFormat(this.{{$c.Name|lowerName}},"yyyy-MM-dd hh:mm:ss")
			{{- else if or ($c.Con|qCon|DP) (and (not ($c.Con|qCon|DTP)) ($c.Con|DP))  }}
			this.queryData.{{$c.Name}} = this.$utility.dateFormat(this.{{$c.Name|lowerName}},"yyyy-MM-dd")
			{{- end -}}
      {{- end}}
      let res = this.$http.xpost("/{{.Name|rmhd|rpath}}/query",this.queryData)
			this.dataList.items = res.items
			this.dataList.count = res.count
    },
    /**改变页容量*/
		pageSizeChange(val) {
      this.paging.ps = val
      this.query()
    },
    /**改变当前页码*/
    pageIndexChange(val) {
      this.paging.pi = val
      this.query()
    },
    /**重置添加表单*/
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
		},
		{{- if gt ($rows|detail|len) 0}}
		showDetail(val){
			var data = {
        {{range $i,$c:=$pks}}{{$c}}: val.{{$c}},{{end}}
      }
      this.$emit("addTab","详情"+val.{{range $i,$c:=$pks}}{{$c}}{{end}},"/{{.Name|rmhd|rpath}}/detail",data);
		},
		{{- end}}
		{{- if gt ($rows|create|len) 0}}
    showAdd(){
      this.$refs.Add.show();
		},
		{{- end}}
		{{- if gt ($rows|update|len) 0}}
    showEdit(val) {
      this.$refs.Edit.editData = val
      this.$refs.Edit.show();
		},
		{{- end}}
		{{- if gt ($rows|delete|len) 0}}
    del(val){
			this.$confirm("此操作将永久删除该数据, 是否继续?", "提示", {confirmButtonText: "确定",  cancelButtonText: "取消", type: "warning"})
			.then(() => {
				this.$http.del("/{{.Name|rmhd|rpath}}", {data:val}, {}, true, true)
				.then(res => {			
					this.dialogFormVisible = false;
					this.refresh()
				})
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消删除"
        });          
      });
		}
		{{- end}}
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
`
