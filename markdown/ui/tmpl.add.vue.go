package ui

//TmplCreateVue 添加创建弹框页面
const TmplCreateVue = `
{{- $empty := "" -}}
{{- $rows := .Rows -}}
{{- $tb :=. -}}
<template>
  <!-- Add Form -->
  <el-dialog title="添加{{.Desc}}" {{- if gt ($rows|create|len) 5}} width="65%" {{else}} width="25%" {{- end}} :visible.sync="dialogAddVisible">
    <el-form :model="addData" {{if gt ($rows|create|len) 5 -}}:inline="true"{{- end}} :rules="rules" ref="addForm" label-width="110px">
    	{{- range $i,$c:=$rows|create }}
      {{if $c.Con|TA -}}
			<el-form-item label="{{$c.Desc|shortName}}" prop="{{$c.Name}}">
				<el-input type="textarea" :rows="2" placeholder="请输入{{$c.Desc|shortName}}" v-model="addData.{{$c.Name}}">
        </el-input>
			</el-form-item>
			{{- else if $c.Con|RD }}
			<el-form-item  label="{{$c.Desc|shortName}}:" prop="{{$c.Name}}">
				<el-radio-group v-model="addData.{{$c.Name}}" style="margin-left:5px">
        	<el-radio v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :label="item.value">{{"{{item.name}}"}}</el-radio>
				</el-radio-group>
			</el-form-item>
			{{- else if $c.Con|SL }}
			<el-form-item label="{{$c.Desc|shortName}}:" prop="{{$c.Name}}">
				<el-select  placeholder="---请选择---" clearable v-model="addData.{{$c.Name}}" style="width: 100%;">
					<el-option v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
			{{- else if $c.Con|CB }}
			<el-form-item label="{{$c.Desc|shortName}}:" prop="{{$c.Name}}"> 
				<el-checkbox-group v-model="addData.{{$c.Name}}">
					<el-checkbox v-for="(item, index) in {{$c.Name|lowerName}}" :key="index" :value="item.value" :label="item.name"></el-checkbox>
				</el-checkbox-group>
			</el-form-item>
			{{- else if $c.Con|DTP }}
			<el-form-item prop="{{$c.Name}}" label="{{$c.Desc|shortName}}:">
      	<el-date-picker class="input-cos" v-model="addData.{{$c.Name}}" type="datetime" value-format="yyyy-MM-dd HH:mm:ss"  placeholder="选择日期"></el-date-picker>
			</el-form-item>
			{{- else if $c.Con|DP }}
			<el-form-item prop="{{$c.Name}}" label="{{$c.Desc|shortName}}:">
      	<el-date-picker class="input-cos" v-model="addData.{{$c.Name}}" type="date" value-format="yyyy-MM-dd"  placeholder="选择日期"></el-date-picker>
      </el-form-item>
      {{- else -}}
      <el-form-item label="{{$c.Desc|shortName}}" prop="{{$c.Name}}">
				<el-input maxlength="{{$c.Len}}" clearable v-model="addData.{{$c.Name}}" placeholder="请输入{{$c.Desc|shortName}}">
				</el-input>
      </el-form-item>
      {{- end}}
      {{end}}
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
  <!--Add Form -->
</template>

<script>
export default {
	data() {
		return {
			addData: {},
			dialogAddVisible: false,
			{{- range $i,$c:=$rows|create -}}
			{{if or ($c.Con|SL) ($c.Con|CB) ($c.Con|RD) }}
      {{$c.Name|lowerName}}: this.$enum.get("{{(or (dicType $c.Con $tb) $c.Name)|lower}}"),
      {{- end}}
			{{- end}}
			rules: {                    //数据验证规则
				{{- range $i,$c:=$rows|create -}}
				{{if ne ($c|isNull) $empty}}
				{{$c.Name}}: [{ required: true, message: "请输入{{$c.Desc|shortName}}", trigger: "blur" }],
				{{- end}}
				{{- end}}
			},
		}
	},
	props: {
		refresh: {
			type: Function,
				default: () => {
				},
		}
	},
	created(){
	},
	methods: {
		closed() {
			this.refresh()
		},
		resetForm(formName) {
			this.dialogAddVisible = false;
			this.$refs[formName].resetFields();
		},
		show(){
			this.dialogAddVisible = true;
		},
		add(formName) {
			{{- range $i,$c:=$rows|create -}}
			{{- if or ($c.Con|cCon|DTP) (and (not ($c.Con|cCon|DP)) ($c.Con|DTP))}}
			this.addData.{{$c.Name}} = this.$utility.dateFormat(this.addData.{{$c.Name}},"yyyy-MM-dd hh:mm:ss")
			{{- else if or ($c.Con|cCon|DP) (and (not ($c.Con|cCon|DTP)) ($c.Con|DP))}}
			this.addData.{{$c.Name}} = this.$utility.dateFormat(this.addData.{{$c.Name}},"yyyy-MM-dd")
			{{- end -}}
			{{- end}}
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/{{.Name|rmhd|rpath}}", this.addData, {}, true, true)
						.then(res => {
							this.$refs[formName].resetFields()
							this.dialogAddVisible = false
							this.refresh()
						})
						.catch(err => {
							this.$message({
								type: "error",
								message: err.response.data
							});
						})
				} else {
						console.log("error submit!!");
						return false;
				}
			});
		},
	}

}
</script>

<style scoped>
</style>
`
