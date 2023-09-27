document.getElementById('searchBar').addEventListener('input', function() {
    let query = this.value.toLowerCase();
    let tableRows = document.querySelectorAll('#artistTable .searchable');
    let suggestions = document.getElementById('suggestions');
    suggestions.innerHTML = ''; // clear existing suggestions
    
    tableRows.forEach(row => {
        if (query === '') {
            row.style.display = ''; // show all rows when the query is empty
            return;
        }
        
        let isMatch = Array.from(row.cells).some(cell => {
            let cellText = cell.textContent.toLowerCase();
            if (cellText.includes(query)) {
                let readableText = cell.textContent.trim().replace(/_/g, ', '); 
                
                let suggestionDiv = document.createElement('div');
                suggestionDiv.textContent = readableText; // use readableText here
                suggestionDiv.addEventListener('click', function() {
                    document.getElementById('searchBar').value = this.textContent;
                    suggestions.innerHTML = ''; // clear suggestions
                    filterRows(this.textContent.toLowerCase());
                });
                suggestions.appendChild(suggestionDiv);
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
