
# POC - Reporting on Pulse

Used libraries:
- For PDF generation:
    - https://wkhtmltopdf.org/usage/wkhtmltopdf.txt
    - https://github.com/SebastiaanKlippert/go-wkhtmltopdf
    - https://github.com/SebastiaanKlippert/go-wkhtmltopdf-lambda
- For HTML templates:
    - https://masterminds.github.io/sprig/strings.html

<br>

Example Endpoint call: `GET /reports/template1?siteId=101&from=2024-09-01&to=2024-09-30&sections=1,2,3`
