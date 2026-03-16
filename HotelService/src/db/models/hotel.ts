import {InferAttributes, InferCreationAttributes, CreationOptional, Model, DataTypes} from 'sequelize';

import sequelize from './sequelize';

class Hotel extends Model<InferAttributes<Hotel>, InferCreationAttributes<Hotel>> {
    declare id: CreationOptional<number>;
    declare name: string;
    declare address: string;
    declare location: string;
    declare createdAt: CreationOptional<Date>;
    declare updatedAt: CreationOptional<Date>;
    declare rating: number;
    declare ratingCount: number;
}



Hotel.init({
    id: {
        type: DataTypes.INTEGER,
        primaryKey: true,
        autoIncrement: true
    },
    name: {
        type: DataTypes.STRING,
        allowNull: false
    },
    address: {
        type: DataTypes.STRING,
        allowNull: false
    },
    location: {
        type: DataTypes.STRING,
        allowNull: false
    },
    createdAt: {
        type: DataTypes.DATE,
        defaultValue: new Date(),
    },
    updatedAt: {
        type: DataTypes.DATE,
        defaultValue: new Date(),
    },
    rating: {
        type: DataTypes.FLOAT,
        defaultValue: null
    },
    ratingCount: {
        type: DataTypes.INTEGER,
        defaultValue: null
    }
}, {
    sequelize,
    tableName: 'hotels',
    timestamps: true,
    underscored: true
});

export default Hotel;