document.getElementById('searchBar').addEventListener('input', function() {
    let query = this.value.toLowerCase();
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    
    tableRows.forEach(row => {
        if (query === '') {
            row.style.display = ''; // show all rows when the query is empty
            return;
        }
        
        let isMatch = Array.from(row.cells).some(cell => {
            let cellText = cell.textContent.toLowerCase();
            if (cellText.includes(query)) {
                return true;
            }
            return false;
        });
        row.style.display = isMatch ? '' : 'none'; // show row if there is a match, else hide it
    });
});

function filterRows(query) {
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    tableRows.forEach(row => {
        row.style.display = Array.from(row.cells).some(cell => cell.textContent.toLowerCase().includes(query)) ? '' : 'none';
    });
}
