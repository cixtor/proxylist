#!/usr/bin/env node
/* jshint esversion: 6 */
/* jshint loopfunc: true */

try {
    let filesystem = require('fs');
    let cheerio = require('cheerio');
    let fpath = null;

    var extractProxyData = function (fpath, callback) {
        filesystem.readFile(fpath, function (err, data) {
            if (err !== null) {
                throw err;
            }

            let buf = new Buffer.from(data.toString(), 'base64').toString();
            let $ = cheerio.load(buf);

            let content = $('td').eq(1).html();
            let pattern = '\\.([0-9a-zA-Z]{4})\{display:(inline|none)\}';
            let cssclass = content.match(new RegExp(pattern, 'g'));
            let invisible = [];
            let parts = [];

            for (var key in cssclass) {
                if (cssclass.hasOwnProperty(key)) {
                    parts = cssclass[key].match(new RegExp(pattern));
                    if (parts.length && parts[2] === 'none') {
                        invisible.push(parts[1]);
                    }
                }
            }


            content = content.replace(/<(span|div) style="display:none">[^<]+<\/(span|div)>/, '');

            for (var idx in invisible) {
                if (invisible.hasOwnProperty(idx)) {
                    pattern = '<(span|div) class="' + invisible[idx] + '">[^<]+<\\/(span|div)>';
                    content = content.replace(new RegExp(pattern), '');
                }
            }

            content = content.replace(/<(span|div) style="display: inline">([^<]+)<\/(span|div)>/g, '$2');
            content = content.replace(/<(span|div) class="[^"]+">([^<]+)<\/(span|div)>/g, '$2');
            content = content.replace(/<\/style>([0-9\.]{7,15}).*/, '$1');

            callback({
                // TableRow: $('tr').attr('rel'),
                LastUpdate: $('span.updatets').text().trim(),
                Address: content.trim().split('\n').pop(),
                Port: $('td').eq(2).text().trim(),
                Country: $('td').eq(3).text().trim(),
                Speed: $('.response_time .indicator').css('width'),
                Connection: $('.connection_time .indicator').css('width'),
                Protocol: $('td').eq(6).text().trim(),
                Anonimity: $('td').eq(7).text().trim(),
            });
        });
    };

    for (var key in process.argv) {
        if (process.argv.hasOwnProperty(key)) {
            fpath = process.argv[key];

            if (fpath.indexOf('tests/') === 0) {
                extractProxyData(fpath, function (proxy) {
                    console.log(fpath, JSON.stringify(proxy, null, '\t'));
                });
            }
        }
    }
} catch (e) {
    console.log('Unexpected error during execution');
    console.log('Be sure to install "cheerio"');
    console.log('$ npm install cheerio');
    console.log();
    console.log(e);
}
