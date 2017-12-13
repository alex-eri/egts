package egts_protocol_data_types

type EGTS_SR_RECORD_RESPONSE struct {
	// Подзапись применяется для осуществления подтверждения процесса
	// обработки записи Протокола Уровня Поддержки Услуг.
	// Данный тип подзаписи должен поддерживаться всеми Сервисами.
	CRN int
	RST byte
	*ServiceDataSubrecord
}

type EGTS_SR_SERVICE_PART_DATA struct {
	// Формат заголовка передаваемой сущности подзаписи


	// характеристика принадлежности передаваемой сущности
	OA byte

	// тип сущности по содержанию. Определены следующие значения данного поля:
	// 00 = данные внутреннего ПО («прошивка»)
	// 01 = блок конфигурационных параметров
	OT byte

	// тип модуля, для которого предназначена передаваемая сущность.
	// Определены следующие значения данного поля:
	// 00 = периферийное оборудование
	// 01 = АТ
	MT byte

	// номер компонента в случае принадлежности сущности непосредственно АТ
	// или идентификатор периферийного модуля/порта, подключенного к АТ,
	// в зависимости от значения параметра MT
	CMI byte

	// версия передаваемой сущности
	// до точки – major version, младший,
	// (старший байт

		// точки – minor version, например версия 2.34 будет представлена числом 0x0222))
		//
	VER int

	// сигнатура (контрольная сумма), всей передаваемой сущности.
	// Используется алгоритм СRC16-CCITT
	WOS int

	// имя файла передаваемой сущности (данное поле опционально и
	// может иметь нулевую длину)
	FN string

	// разделитель строковых параметров (всегда имеет значение 0)
	D byte

}

type EGTS_SR_SERVICE_PART_DATA struct {
	// Подзапись предназначена для передачи на АТ данных, которые разбиваются
	// на части и передаются последовательно. Данная подзапись применяется
	// для передачи больших объектов, длина которых не позволяет передать их
	// на АТ одним пакетом.

	// уникальный идентификатор передаваемой сущности. Инкрементируется при
	// начале отправки новой сущности. Данный параметр позволяет однозначно
	// идентифицировать, какой именно сущности данная часть принадлежит.
	ID int

	// последовательный номер текущей части передаваемой сущности
	PN int

	// ожидаемое количество частей передаваемой сущности
	EPQ int

	// заголовок, содержащий параметры, характеризующие передаваемую сущность.
	// Данный заголовок передаётся только для первой части сущности.
	// При передаче второй и последующих частей, данное поле не передается.
	// Структура заголовка ODH представлена в Таблице 25.
	ODH []EGTS_SR_SERVICE_PART_DATA

	// непосредственно данные передаваемой сущности
	OD string


}

type EGTS_SR_SERVICE_FULL_DATA struct {
	// заголовок, содержащий параметры, характеризующие передаваемую сущность.
	// Для подзаписи EGTS_SR_SERVICE_FULL_DATA параметр ODH является
	// обязательным и присутствует в каждой такой подзаписи.
	ODH string

	// непосредственно данные передаваемой сущности
	OD string
}

type EGTS_FIRMWARE_SERVICE struct {
	// Данный тип сервиса предназначен для передачи на АТ конфигурации и
	// обновления ПО аппаратной части модулей и блоков самого АТ,
	// а также периферийного оборудования, подключенного к АТ.

	EGTS_FIRMWARE_SERVICE []EGTS_FIRMWARE_SERVICE
	EGTS_SR_SERVICE_PART_DATA []EGTS_SR_SERVICE_PART_DATA
	EGTS_SR_SERVICE_FULL_DATA []EGTS_SR_SERVICE_FULL_DATA
}