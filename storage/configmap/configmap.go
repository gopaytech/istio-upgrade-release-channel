package configmap

// import (
// 	"context"
// 	"log"

// 	"github.com/gopaytech/istio-upgrade-release-channel/config"
// 	"github.com/gopaytech/istio-upgrade-release-channel/settings"
// 	"github.com/gopaytech/istio-upgrade-release-channel/types"
// 	v1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// )

// func NewConfigmapStorage(kubernetesConfig *config.Kubernetes, settings settings.Settings) ConfigmapStorage {
// 	return ConfigmapStorage{
// 		KubernetesConfig: kubernetesConfig,
// 		Settings:         settings,
// 	}
// }

// type ConfigmapStorage struct {
// 	KubernetesConfig *config.Kubernetes
// 	Settings         settings.Settings
// }

// func (s ConfigmapStorage) Create(ctx context.Context, upgrade types.Upgrade) error {
// 	c := s.KubernetesConfig.Client()

// 	cm := v1.ConfigMap{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      s.Settings.StorageConfigMapName,
// 			Namespace: s.Settings.StorageConfigMapNameSpace,
// 		},
// 		Data: upgrade.ToConfigMapData(),
// 	}

// 	_, err := c.CoreV1().ConfigMaps(s.Settings.StorageConfigMapNameSpace).Create(ctx, &cm, metav1.CreateOptions{})
// 	if err != nil {
// 		log.Println("failed to create configmap: ", err.Error())
// 		return err
// 	}

// 	return nil
// }

// func (s ConfigmapStorage) Get(ctx context.Context) (*types.Upgrade, error) {
// 	c := s.KubernetesConfig.Client()

// 	_, err := c.CoreV1().ConfigMaps(s.Settings.StorageConfigMapNameSpace).Get(ctx, s.Settings.StorageConfigMapName, metav1.GetOptions{})
// 	if err != nil {
// 		log.Printf("failed to get configmap %s %s\n", s.Settings.StorageConfigMapName, err.Error())
// 		return nil, err
// 	}

// 	// todo: configmap to types.upgrade
// 	return &types.Upgrade{}, nil
// }

// func (s ConfigmapStorage) Update(ctx context.Context, upgrade types.Upgrade) error {
// 	c := s.KubernetesConfig.Client()

// 	cm, err := c.CoreV1().ConfigMaps(s.Settings.StorageConfigMapNameSpace).Get(ctx, s.Settings.StorageConfigMapName, metav1.GetOptions{})
// 	if err != nil {
// 		log.Printf("failed to get configmap %s %s\n", s.Settings.StorageConfigMapName, err.Error())
// 		return err
// 	}

// 	cm.Data = upgrade.ToConfigMapData()

// 	_, err = c.CoreV1().ConfigMaps(s.Settings.StorageConfigMapNameSpace).Update(ctx, cm, metav1.UpdateOptions{})
// 	if err != nil {
// 		log.Printf("failed to update configmap %s %s\n", s.Settings.StorageConfigMapName, err.Error())
// 		return err
// 	}

// 	return nil
// }
