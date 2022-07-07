package builder

/*
当构建的对象很大并且需要多个步骤时，请使用构建器模式。它有助于减少构造函数的大小。房屋的建造变得简单，不需要大型建筑
当需要创建同一产品的不同版本时。例如，在下面的代码中，我们看到了不同版本的house ie。冰屋和由iglooBuilder和normalBuilder建造的普通房屋
当半构造的最终对象不应该存在时。再次引用下面的代码，创建的房屋将被完全创建或根本不创建。混凝土建造者结构持有正在创建的房屋对象的临时状态
*/

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
