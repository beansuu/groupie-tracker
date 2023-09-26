document.getElementById('searchBar').addEventListener('input', function() {
    let query = this.value.toLowerCase();
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    
    tableRows.forEach(row => {
        row.style.display = row.textContent.toLowerCase().includes(query) ? '' : 'none';
    });
});