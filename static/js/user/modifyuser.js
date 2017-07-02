$(function () {
  // modifyAdmUserObj = {
  //       search: function () {
  //           $('#updateUserGroup').datagrid('load', {
  //               groupName: $('input[name="modifyUser_UserGroupName"]').val()         
  //           });
  //       }
  //   }
  //   //datagrid初始化
  //   $('#updateUserGroup').datagrid({
  //       url: 'user/gridgrouplist',
  //       queryParams: { userId: $("input[name='userId']").val() },
  //       iconCls: 'icon-edit',//图标
  //       width: 700,
  //       height: 'auto',
  //       nowrap: false,
  //       striped: true,
  //       border: true,
  //       collapsible: false,//是否可折叠的
  //       fit: true,//自动大小
  //       //sortName: 'code',
  //       //sortOrder: 'desc',
  //       remoteSort: false,
  //       idField: 'id',
  //       singleSelect: false,//是否单选
  //       pagination: true,//分页控件
  //       rownumbers: true,//行号
  //       fitColumns: true,//列宽自适应（列设置width=100）
  //       frozenColumns: [[
  //           { field: 'ck', checkbox: true }
  //       ]],//设置表单复选框
  //       toolbar: modifyUser_toolbar,
  //       onLoadSuccess:function(row){//当表格成功加载时执行               
  //               var rowData = row.rows;
  //               $.each(rowData,function(idx,val){//遍历JSON
  //                     if(val.check==true){
  //                       $("#updateUserGroup").datagrid("selectRow", idx);//如果数据行为已选中则选中改行
  //                     }
  //               });              
  //           }
  //   });
  $("#getUserCatogery").combobox({          
        url: 'usercatogery/listall',
        valueField:'catogeryname',//相当于option的value值
        textField:'catogeryname',//相当于<option></option>之间的显示值 value:1000    //默认显示值
    });
})


function submitModifyUserForm() {
    // var selections = $('#updateUserGroup').datagrid('getSelections')
    // if (selections.length == 0) {
    //     $.messager.alert('操作提示', "请至少选择一个组", 'info');
    //     return false
    // }

    // var idArray = new Array(selections.length)
    // for (var i = 0; i < selections.length; i++) {
    //     idArray[i] = selections[i].id
    // }
    // ids = idArray.join(",")

    url = "/user/modifyyuser"
    var data = {
        userId:$("input[name='admUserId']").val(),
        account: $("input[name='modifyUserAcout']").val(),
        name: $("input[name='modifyUserName']").val(),
        phone: $("input[name='modifyUserPhone']").val(),
        catogery: $("input[name='modifyUserCatogery']").val(),
        password: $("input[name='modifyUserPassword']").val(),
        mail: $("input[name='modifyUserEmail']").val()
    };

    if (data.account.length < 1 || data.name.length < 1 || data.phone.length < 1 || data.catogery.length < 1 || data.mail.length < 1) {
        $.messager.alert('操作提示', "信息填写不完整,请补充后重新提交", 'info');
        return
    }


    $.post(url, data, function (result) {
        if (result == "success") {
            $('#modifyUser').window("close")
            $.messager.alert('操作提示', "修改成功", 'info');
            loadModifyUserGrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}


function clearModifyUserForm() {
    $('#modifyUser').form('clear');
}


function loadModifyUserGrid() {
    $('#user_list').datagrid('load', {
    });
}