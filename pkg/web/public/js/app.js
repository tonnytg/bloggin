function delete_post(id) {
    if (window.confirm("Are you sure to delete Post ID:" + id)) {

        fetch('/posts?id=' + id, {
            method: 'DELETE',
        })
            .then(res => res.text()) // or res.json()
            .then(res => console.log(res))
    }
    // Do something;
    window.location.reload()
};