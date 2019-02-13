$(document).ready(function(){
    contextLoad()    
})

function contextLoad(){
    $.get( "/api/context", function( data ) {
        var table;

        table = $("#avilableContextTable tbody")
        $.each(data.contexts, (k, v) => {

            // Get Cluster
            var cluster = data.clusters.find(x => x.name === v.context.cluster) || "N/A"
            var user = data.users.find(x => x.name === v.context.user) || "N/A"

            var row = "<tr>";

            if (data.currentContext == v.name) {
                row = '<tr class="table-success">';
            }

            row += "<td>" + v.name + "</td>";
            row += '<td><span class="clusterDialog">' + cluster.name + '</span></td>';
            row += '<td><span class="userDialog">' + user.name + '</span></td>';
            row += "<td>";
            row += '<button type="button" class="btn btn-outline-secondary"><i class="far fa-edit"></i></button>';
            row += '<button type="button" class="btn btn-outline-danger"><i class="far fa-trash-alt"></i></button>';
            row += '<button type="button" class="btn btn-outline-success"><i class="fas fa-check"></i></button>';
            row += "</td>";
            row += "</tr>";
            table.append(row) 
            
            $(".clusterDialog").on("click", function(e){
                alert($(event.target).text())
                console.log(e)
            })
            
            $(".userDialog").on("click", function(e){
                alert($(event.target).text())
                console.log(e)
            })
        })
    });
}

