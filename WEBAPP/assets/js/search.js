$('#searching-form').on('submit',startSearching)

function startSearching(event){
    //event.preventDefault();
    console.log("searching")

    $.ajax({
        url: "/",
        method: "POST",
        data: {
            path: $('#path').val(),
            word: $('#word').val(),
        }
    })
}