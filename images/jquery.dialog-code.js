var gitbook = gitbook || [];
gitbook.push(function() {
    document.querySelectorAll('.DialogCode').forEach(el => {
        getSourceCode($(el).data('code'), el)
    })
});

function getSourceCode(word, parent) {
    let keywords = [
        'newproc', 'newproc1', 'runqput', 'wakep', 'startm',
        'mainPC',
    ];
    $.get(`../docs/go1.16.10/${word}.html`, function (text) {
        let newElem = $(text);
        let spans = newElem.find('span');
        spans.each(function (k, span) {
            if (keywords.includes(span.innerText) && word != span.innerText) {
                spans.eq(k).html(`<a onclick="getSourceCode('${span.innerText}');">${span.innerText}</a>`)
            }
        });
        if (parent) {
            $(parent).append(newElem)
        } else {
            $('body').append(newElem);
            newElem.dialog({
                'title': word
            }, function (api) {
                api.open();
            });
        }
    })
}