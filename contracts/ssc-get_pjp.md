**GET /api/ssc/v1/pjp?page=1&length=2**
----
  Returns array of permanent journey plan dataset.
* **URL Params**  
  Optional: page=[integer]  
  Optional: length=[integer]
* **Data Params**  
  None
* **Headers**  
  Content-Type: application/json  
  Authorization: Bearer `<JWT Token>`
* **Success Response:**  
* **Code:** 200  
  **Content:**  
```
{
  "status": 200,
  "type": null,
  "data": [
        {
          "id": 1,
          "createdAt": "0001-01-01T00:00:00Z",
          "updatedAt": "0001-01-01T00:00:00Z",
          "deletedAt": null,
          "supplierId": 1,
          "portfolioId": 1,
          "storeId": "1",
          "storeExternalId": "1",
          "storeName": "TOKO TEST 1",
          "salesRepId": "1",
          "salesRepExternalId": "1",
          "salesRepName": "SALES REP TEST 1",
          "isRepetitive": false,
          "repeatEveryValue": null,
          "repeatEveryMeasure": null,
          "repeatUntil": "2021-04-01T06:59:59+07:00",
          "route": 1,
          "visitSchedules": [
            {
                "id": 1,
                "createdAt": "0001-01-01T00:00:00Z",
                "updatedAt": "0001-01-01T00:00:00Z",
                "deletedAt": null,
                "permanentJourneyPlanID": 1,
                "visitDate": "2021-03-02T06:59:59+07:00"
            },
            {
                "id": 2,
                "createdAt": "0001-01-01T00:00:00Z",
                "updatedAt": "0001-01-01T00:00:00Z",
                "deletedAt": null,
                "permanentJourneyPlanID": 1,
                "visitDate": "2021-03-03T06:59:59+07:00"
            },
            {
                "id": 3,
                "createdAt": "0001-01-01T00:00:00Z",
                "updatedAt": "0001-01-01T00:00:00Z",
                "deletedAt": null,
                "permanentJourneyPlanID": 1,
                "visitDate": "2021-03-04T06:59:59+07:00"
            },
            {
                "id": 4,
                "createdAt": "0001-01-01T00:00:00Z",
                "updatedAt": "0001-01-01T00:00:00Z",
                "deletedAt": null,
                "permanentJourneyPlanID": 1,
                "visitDate": "2021-03-05T06:59:59+07:00"
            }
          ]
        },
        {
          "id": 2,
          "createdAt": "0001-01-01T00:00:00Z",
          "updatedAt": "0001-01-01T00:00:00Z",
          "deletedAt": null,
          "supplierId": 1,
          "portfolioId": 1,
          "storeId": "1",
          "storeExternalId": "1",
          "storeName": "TOKO TEST 2",
          "salesRepId": "1",
          "salesRepExternalId": "1",
          "salesRepName": "SALES REP TEST 2",
          "isRepetitive": false,
          "repeatEveryValue": null,
          "repeatEveryMeasure": null,
          "repeatUntil": "2021-04-01T06:59:59+07:00",
          "route": 1,
          "visitSchedules": []
        }
  ],
  "message": null,
  "meta": {
    "total": 3,
    "limit": 2,
    "skip": 1
  },
  "code": null
}
```