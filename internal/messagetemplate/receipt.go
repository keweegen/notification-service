package messagetemplate

import (
    "github.com/volatiletech/sqlboiler/v4/types"
    "html/template"
)

type ReceiptTemplate struct {
    OrderID          int    `json:"orderId"`
    CommissionAmount string `json:"commissionAmount"`
    TotalAmount      string `json:"totalAmount"`
}

func (r *ReceiptTemplate) Name() string {
    return Receipt.String()
}

func (r *ReceiptTemplate) SetParams(data types.JSON) error {
    return data.Unmarshal(&r)
}

func (r *ReceiptTemplate) EmailTemplate() *template.Template {
    return receiptEmailTemplate
}

func (r *ReceiptTemplate) TelegramTemplate() *template.Template {
    return receiptTelegramTemplate
}

var receiptEmailTemplate = template.Must(template.New("ns.email.receipt").Parse(`<h3>Чек</h3>

<p>Заказ <b>{{.OrderID}}</b> успешно оплачен</p>

<p>
    Комиссия: {{.CommissionAmount}} <br/>
    Сумма к списанию: {{.TotalAmount}}
</p>

<p>Спасибо за покупку!</p>`))

var receiptTelegramTemplate = template.Must(template.New("ns.telegram.receipt").Parse(`<b>Чек</b>

Заказ <code>{{.OrderID}}</code> успешно оплачен

Комиссия: {{.CommissionAmount}}
Сумма к списанию: {{.TotalAmount}}

Спасибо за покупку`))
