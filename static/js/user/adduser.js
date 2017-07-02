
$(function () {
    // addUserObj = {
    //     search: function () {
    //         $('#usergroup').datagrid('load', {
    //             groupName: $('input[name="addUser_UserGroupName"]').val()         
    //         });
    //         $('#getUserCatogery').datagrid('load', {
    //             groupName: $('input[name="addUser_UserGroupName"]').val()         
    //         });
    //     }
    // }
    // var userid = $('#getUserid').combobox('getValue');
    //datagrid初始化
    // $('#usergroup').datagrid({
    //     url: 'usercatogery/gridlist',
    //     // queryParams: { roleid: 0 },
    //     iconCls: 'icon-edit',//图标
    //     width: 700,
    //     height: 'auto',
    //     nowrap: false,
    //     striped: true,
    //     border: true,
    //     collapsible: false,//是否可折叠的
    //     fit: true,//自动大小
    //     //sortName: 'code',
    //     //sortOrder: 'desc',
    //     remoteSort: false,
    //     idField: 'id',
    //     singleSelect: true,//是否单选
    //     pagination: true,//分页控件
    //     rownumbers: true,//行号
    //     fitColumns: true,//列宽自适应（列设置width=100）
    //     frozenColumns: [[
    //         { field: 'ck', checkbox: true }
    //     ]],//设置表单复选框
    //     toolbar: addUser_toolbar
    // });

    $("#getUserCatogery").combobox({          
        url: 'usercatogery/listall',
        valueField:'catogeryname',//相当于option的value值
        textField:'catogeryname',//相当于<option></option>之间的显示值 value:1000    //默认显示值
    });
})



function submitAddUserForm() {
    // var selections = $('#usergroup').datagrid('getSelections')
    // if (selections.length == 0) {
    //     $.messager.alert('操作提示', "请至少选择一个组", 'info');
    //     return false
    // }

    // var idArray = new Array(selections.length)
    // for (var i = 0; i < selections.length; i++) {
    //     idArray[i] = selections[i].id
    // }
    // ids = idArray.join(",")

    url = "/user/adduser"
    var data = {
        account: $("input[name='userAcout']").val(),
        name: $("input[name='userName']").val(),
        phone: $("input[name='userPhone']").val(),
        catogery: $("input[name='userCatogery']").val(),
        password: $("input[name='userPassword']").val(),
        mail: $("input[name='userEmail']").val()
    };

    if (data.account.length < 1 || data.name.length < 1 || data.phone.length < 1 || data.catogery.length < 1 || data.password.length < 1 || data.mail.length < 1) {
        $.messager.alert('操作提示', "信息填写不完整,请补充后重新提交", 'info');
        return
    }


    $.post(url, data, function (result) {
        if (result == "success") {
           $('#addUser').window("close")
            $.messager.alert('操作提示', "添加成功", 'info');
           loadUserGrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}


function clearAddUserForm() {
    $('#addUser').form('clear');
}

function loadUserGrid() {
    $('#user_list').datagrid('load', {
    });
}