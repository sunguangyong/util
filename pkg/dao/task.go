package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type Task struct {
	ID                           int64   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Code                         string  `gorm:"column:code;NOT NULL" json:"code"`
	TaskNo                       string  `gorm:"column:task_no;NOT NULL" json:"task_no"`
	Name                         string  `gorm:"column:name;NOT NULL" json:"name"`
	CreateType                   int64   `gorm:"column:create_type;default:0" json:"create_type"`
	Mode                         int64   `gorm:"column:mode;default:0;NOT NULL" json:"mode"`
	Scene                        int64   `gorm:"column:scene;default:0;NOT NULL" json:"scene"`
	TaskMileage                  float64 `gorm:"column:task_mileage;default:0;NOT NULL" json:"task_mileage"`
	CreatorID                    int64   `gorm:"column:creator_id;default:0" json:"creator_id"`
	CreatorName                  string  `gorm:"column:creator_name" json:"creator_name"`
	UserName                     string  `gorm:"column:user_name" json:"user_name"`
	MediaBucket                  string  `gorm:"column:media_bucket;NOT NULL" json:"media_bucket"`
	Source                       int64   `gorm:"column:source;default:0;NOT NULL" json:"source"`
	Status                       int64   `gorm:"column:status;default:0;NOT NULL" json:"status"`
	ActiveTime                   string  `gorm:"column:active_time;NOT NULL" json:"active_time"`
	ImgNum                       int64   `gorm:"column:img_num;default:0" json:"img_num"`
	TaskExportID                 int64   `gorm:"column:task_export_id;default:0" json:"task_export_id"`
	ProjectCode                  string  `gorm:"column:project_code;NOT NULL" json:"project_code"` // 项目编码
	DronePilot                   string  `gorm:"column:drone_pilot;NOT NULL" json:"drone_pilot"`   // 飞手
	CreatedAt                    int64   `gorm:"column:created_at;default:0;NOT NULL" json:"created_at"`
	UpdatedAt                    int64   `gorm:"column:updated_at;default:0;NOT NULL" json:"updated_at"`
	DeletedAt                    int64   `gorm:"column:deleted_at;default:0;NOT NULL" json:"deleted_at"`
	DroneID                      int64   `gorm:"column:drone_id;default:0" json:"drone_id"`
	LineID                       int64   `gorm:"column:line_id;default:0" json:"line_id"`
	MainLineID                   int64   `gorm:"column:main_line_id;default:0" json:"main_line_id"`
	OrgID                        int64   `gorm:"column:org_id;default:0" json:"org_id"`
	UserID                       int64   `gorm:"column:user_id;default:0" json:"user_id"`
	OrderID                      int64   `gorm:"column:order_id;default:0" json:"order_id"`
	InspectionCompanyID          int64   `gorm:"column:inspection_company_id;default:0;NOT NULL" json:"inspection_company_id"`
	DirectorID                   int64   `gorm:"column:director_id;default:0;NOT NULL" json:"director_id"` // id
	FinishTimeStr                string  `gorm:"column:finish_time_str;NOT NULL" json:"finish_time_str"`
	Remark                       string  `gorm:"column:remark;NOT NULL" json:"remark"` // 备注，xuelian中用于写任务的终止原因
	AppVisible                   int     `gorm:"column:app_visible;default:0;NOT NULL" json:"app_visible"`
	ExecMethod                   int     `gorm:"column:exec_method;default:1;NOT NULL" json:"exec_method"`                                           // 任务执行方式 1:APP 2:机场
	AirportID                    int64   `gorm:"column:airport_id;default:0;NOT NULL" json:"airport_id"`                                             // 机场id
	StartTime                    int64   `gorm:"column:start_time;default:0" json:"start_time"`                                                      // 计划任务开始时间
	EndTime                      int64   `gorm:"column:end_time;default:0" json:"end_time"`                                                          // 计划任务结束时间
	AuditStatus                  int64   `gorm:"column:audit_status;default:0" json:"audit_status"`                                                  // 审核状态 0 未审核 1 审核中 2 待提交 3 复核中 4 已完成
	TaskType                     int64   `gorm:"column:task_type;default:0" json:"task_type"`                                                        // 任务类型（1:平台端上传航迹创建的任务）
	RoutedocsCompressFileMediaID int64   `gorm:"column:routedocs_compress_file_media_id;default:0;NOT NULL" json:"routedocs_compress_file_media_id"` // 任务航迹文件的压缩文件
	AuditDeviceStatusID          int64   `gorm:"column:audit_device_status_id;default:0" json:"audit_device_status_id"`                              // 设备审核状态 0 未审核 1 审核中 2 已完成
	Category                     int     `gorm:"column:category;default:0;NOT NULL" json:"category"`                                                 // 任务类别 （0首飞 1复飞第1次，n复飞第n次）
	JijianStatus                 int     `gorm:"column:jijian_status;default:0;NOT NULL" json:"jijian_status"`                                       // 马兰基建任务状态（0.未开始 1.执行中 2.已完成）
	AiDeviceStatusID             int64   `gorm:"column:ai_device_status_id;default:0" json:"ai_device_status_id"`                                    // 算法检测状态 0 未检测 1 已检测
	ExecOrgID                    int64   `gorm:"column:exec_org_id;default:0" json:"exec_org_id"`                                                    // 执行任务的组织单位id
	InspectReason                string  `gorm:"column:inspect_reason" json:"inspect_reason"`                                                        // 巡视原因
}

type TaskModel struct {
	db *gorm.DB
}

func NewTaskModel(db *gorm.DB) *TaskModel {
	return &TaskModel{
		db: db,
	}
}

func (m *TaskModel) TableName() string {
	return "task"
}

func (m *TaskModel) CommonQuery(ctx context.Context, limit, offset int, opts ...Option) (tasks []*Task, err error) {
	tasks = make([]*Task, 0)
	db := m.db.WithContext(ctx).Model(&Task{})
	for _, opt := range opts {
		db = opt(db)
	}

	if limit > -1 {
		db.Limit(limit)
	}

	if offset > -1 {
		db.Offset(offset)
	}

	err = db.Find(&tasks).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return tasks, nil
}

func (m *TaskModel) WithID(id int64) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func (m *TaskModel) WithName(name string) Option {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}
