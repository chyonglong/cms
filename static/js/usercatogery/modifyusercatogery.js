
//设置tree的初始化参数
// var usercatogerysetting = {
//     check: {
//         enable: true
//         // chkboxType: { "Y": "", "N": "" }
//     },
//     data: {
//         simpleData: {
//             enable: true
//         }
//     }
// };

//初始化tree
// // $(document).ready(loadTree());
// function loadTree() {
//     var admgroupuserid = $("input[name='admgroupuserid']").val()
//     url = "/usercatogery/loadtreechecked"
//     var data = {
//         admgroupuserid: admgroupuserid,
//     };
//     $.post(url, data, function (result) {
//         // zNodes = result
//         $.fn.zTree.init($("#modifyadmgrouproletree"), usercatogerysetting, result);
//     });
// }

/**
 * 修改管理员组
 */
function submitModifyUserCatogeryForm() {
    // var zTree = $.fn.zTree.getZTreeObj("modifyadmgrouproletree");
    // nodes = zTree.getCheckedNodes(true);
    // checkCount = nodes.length;
    // //判断选中的节点数，如果没有选中节点则提示操作错误
    // if (checkCount == 0) {
    //     $.messager.alert('操作提示', "请至少选择一个权限", 'info');
    //     return false;
    // }
    // //获取所有选中的节点ID
    // var idArray = new Array(checkCount)
    // for (var i = 0; i < nodes.length; i++) {
    //     idArray[i] = nodes[i].id
    // }
    // ids = idArray.join(",")

    url = "/usercatogery/modifyusercatogery"

    var data = {
        id: $("input[name='usercatogeryid']").val(),
        catogeryname: $("input[name='ag_m_name']").val(),
        describe: $("input[name='ag_m_describe']").val()
    };

    $.post(url, data, function (result) {
        if (result == "success") {
            $('#modifyusercatogery').window("close")
            $.messager.alert('操作提示', "修改成功", 'info');
            loadUserCatogeryDatagrid()
        } else {
            $.messager.alert('操作提示', result, 'info');
        }
    });
}

function clearModifyUserCatogeryForm() {
    $('#modifyusercatogery').form('clear');
}


